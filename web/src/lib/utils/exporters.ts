// src/lib/utils/exporters.ts
import type {
  PostFilterDataRequest,
  PostStaticFireColumnsRequest,
} from "$lib/models/dashboardModels";
import type { SelectedFile } from "$lib/models/selectedFile";

/** Narrow function types so this module is UI-agnostic */
export type FetchStaticFireColumnsFn = (
  req: PostStaticFireColumnsRequest
) => Promise<
  | {
      xColumns: Record<string, { rows: number[] }>;
      yColumns: Record<string, { rows: number[] }>;
    }
  | null
>;

export type FetchStaticFireFilteredDataFn = (
  req: PostFilterDataRequest
) => Promise<
  | {
      xColumns: number[];
      yColumns: number[];
    }
  | null
>;

/** Internal: pick numeric parameter based on selected filter */
function pickFilterValue(
  filterNumber: number,
  sigma: number,
  windowSize: number
): number {
  if (filterNumber === 1) return sigma; // Gaussian
  if (filterNumber === 2) return windowSize; // Moving Avg
  return 0;
}

export function labelForFilter(filterNumber: number): string {
  return filterNumber === 1 ? "Gaussian" : filterNumber === 2 ? "Moving Avg" : "None";
}

/** Fetch raw series for the current x/y selections */
export async function getRawSeries(
  selectedFile: SelectedFile,
  xIndex: number,
  yIndex: number,
  startRow: number,
  numRows: number,
  fetchStaticFireColumns: FetchStaticFireColumnsFn
) {
  const xColumnName = selectedFile.metadata.xColumnNames[xIndex];
  const yColumnName = selectedFile.metadata.yColumnNames[yIndex];

  const baseReq: PostStaticFireColumnsRequest = {
    name: selectedFile.name,
    startRow,
    numRows,
    xColumnNames: [xColumnName],
    yColumnNames: [yColumnName],
  };

  const baseRes = await fetchStaticFireColumns(baseReq);
  if (!baseRes) throw new Error("Could not fetch base series.");

  const xRaw = baseRes.xColumns[xColumnName]?.rows ?? [];
  const yRaw = baseRes.yColumns[yColumnName]?.rows ?? [];

  return { xRaw, yRaw, xColumnName, yColumnName };
}

/** Apply filter remotely using existing backend */
export async function getFilteredSeriesRemote(
  x: number[],
  y: number[],
  filterNumber: number,
  sigma: number,
  windowSize: number,
  fetchStaticFireFilteredData: FetchStaticFireFilteredDataFn
) {
  if (filterNumber === 0) throw new Error("No filter selected.");

  const filterValue = pickFilterValue(filterNumber, sigma, windowSize);
  const req: PostFilterDataRequest = {
    xColumns: x,
    yColumns: y,
    filterValue,
    filterNumber,
  };

  const res = await fetchStaticFireFilteredData(req);
  if (!res?.xColumns || !res?.yColumns) throw new Error("Filter failed.");

  return { x: res.xColumns, y: res.yColumns };
}

/** Build an Excel Blob (lazy-import xlsx to avoid SSR issues) */
export async function makeFilteredExcelBlob(
  x: number[],
  y: number[],
  headerX: string,
  headerY: string,
  filterLabel: string
): Promise<Blob> {
  const XLSX = await import("xlsx"); // dynamic import: client only
  const rows: (string | number | null)[][] = [];
  rows.push([headerX, `${headerY} (${filterLabel})`]);

  const n = Math.max(x.length, y.length);
  for (let i = 0; i < n; i++) {
    rows.push([
      i < x.length ? x[i] : null,
      i < y.length ? y[i] : null,
    ]);
  }

  const ws = XLSX.utils.aoa_to_sheet(rows);
  const wb = XLSX.utils.book_new();
  XLSX.utils.book_append_sheet(wb, ws, "Filtered");
  const wbout = XLSX.write(wb, { bookType: "xlsx", type: "array" });

  return new Blob([wbout], {
    type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
  });
}

/** Download a Blob with a neat timestamped filename */
export function downloadBlob(blob: Blob, baseName: string) {
  const a = document.createElement("a");
  const ts = new Date().toISOString().replace(/[:.]/g, "-");
  a.href = URL.createObjectURL(blob);
  a.download = `${baseName}_${ts}.xlsx`;
  document.body.appendChild(a);
  a.click();
  URL.revokeObjectURL(a.href);
  document.body.removeChild(a);
}

/** One-call “get raw → filter → make xlsx → download” */
export async function exportFilteredExcel(options: {
  selectedFile: SelectedFile;
  xIndex: number;
  yIndex: number;
  startRow: number;
  numRows: number;
  filterNumber: number;
  sigma: number;
  windowSize: number;
  fetchStaticFireColumns: FetchStaticFireColumnsFn;
  fetchStaticFireFilteredData: FetchStaticFireFilteredDataFn;
}) {
  if (typeof window === "undefined") {
    throw new Error("Export must run in the browser.");
  }

  const { xRaw, yRaw, xColumnName, yColumnName } = await getRawSeries(
    options.selectedFile,
    options.xIndex,
    options.yIndex,
    options.startRow,
    options.numRows,
    options.fetchStaticFireColumns
  );
  if (!xRaw.length || !yRaw.length) throw new Error("Empty series.");

  const filtered = await getFilteredSeriesRemote(
    xRaw,
    yRaw,
    options.filterNumber,
    options.sigma,
    options.windowSize,
    options.fetchStaticFireFilteredData
  );

  const filterLabel = labelForFilter(options.filterNumber);
  const blob = await makeFilteredExcelBlob(
    filtered.x,
    filtered.y,
    xColumnName,
    yColumnName,
    filterLabel
  );

  const baseName = `${options.selectedFile.name}_${xColumnName}_${yColumnName}_filtered`;
  downloadBlob(blob, baseName);
}
