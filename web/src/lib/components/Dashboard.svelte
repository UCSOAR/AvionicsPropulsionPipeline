<script lang="ts">
  import Dropdown from "./Dropdown.svelte";
  import Input from "./Input.svelte";
  import IconButton from "./IconButton.svelte";
  import FullscreenButton from "./FullscreenButton.svelte";
  import DropdownBtn from "./DropdownBtn.svelte";

  import { writable, get } from "svelte/store";
  import { onMount } from "svelte";

  import { fetchStaticFireColumns } from "$lib/utils/getStaticFireColumns";
  import { fetchStaticFireFilteredData } from "$lib/utils/getFilteredData";
  import { numericRegex } from "$lib/utils/regexps";
  import { exportFilteredExcel } from "$lib/utils/exporters";
  import { endpointMapping } from "$lib/utils/constants";

  import type { Config, Data, Layout } from "plotly.js";
  import type {
    PostFilterDataRequest,
    PostStaticFireColumnsRequest
  } from "$lib/models/dashboardModels";
  import type { SelectedFile } from "$lib/models/selectedFile";

  import {
    Loader2,
    MessageCircleWarningIcon,
    RefreshCcw,
    ChartLine,
    Download,
    ChevronLeft,
    ChevronRight
  } from "@lucide/svelte";
  import { X } from "lucide-svelte";

  export let selectedFile: SelectedFile;
  export let refreshGraph: () => Promise<void>;

  let plotlyChartDiv: HTMLDivElement;
  let fullscreenTarget: HTMLDivElement;

  const selectedXColumnIndex = writable(0);
  const selectedYColumnIndex = writable(0);

  let selectedFilterIndex = 0;
  let selectedDownloadIndex = 0;

  // User inputs
  let startRow = 0; // user Start Row
  let numRows = 0;  // user Row Count (treated as END index, exclusive)

  // Pagination
  const ROLLING_WINDOW = 250000;
  let isRollingMode = false;
  let pageOffset = 0; // offset from startRow for current page

  // Filter + misc
  let sigma = 0;
  let windowSize = 0;
  let testStart = 0;
  let testEnd = 0;
  let isLoadingPlotly = false;
  let isFilterOn = false;

  let plotError = "";

  const style = {
    margin: 50,
    bgColor: "#1f1f1f",
    txtColor: "#e1e1e1",
    themeColor: "#dc2626"
  };

  const config: Partial<Config> = { responsive: true };
  const shrunkenHeight = 400;

  const layout: Partial<Layout> = {
    autosize: true,
    height: shrunkenHeight,
    margin: {
      l: style.margin,
      r: style.margin,
      t: style.margin,
      b: style.margin
    },
    paper_bgcolor: style.bgColor,
    plot_bgcolor: style.bgColor,
    font: {
      family: "Inter",
      color: "white"
    },
    xaxis: { color: style.txtColor },
    yaxis: { color: style.txtColor },
    legend: {
      orientation: "h",
      x: 0.39
    }
  };

  const safeParseInt = (value: string) => {
    const parsedValue = parseInt(value, 10);
    return isNaN(parsedValue) ? 0 : parsedValue;
  };

  const safeParseFloat = (value: string) => {
    const parsedValue = parseFloat(value);
    return isNaN(parsedValue) ? 0 : parsedValue;
  };

  const handleFilter = () => {
    isFilterOn = !isFilterOn;
    if (data.length > 0) {
      refreshPlotly();
    }
  };

  // Interpret Start Row / Row Count
  // Row Count is treated as an END index (exclusive).
  const getSelectionInfo = () => {
    if (!selectedFile) return null;
    if (numRows <= 0) return null;

    const totalRows = selectedFile.metadata.totalRows ?? 0;

    let start = startRow;
    if (!Number.isFinite(start) || start < 0) start = 0;
    if (start > totalRows) start = totalRows;

    let end = numRows;
    if (!Number.isFinite(end) || end < 0) end = 0;

    if (end < start) end = start;
    if (end > totalRows) end = totalRows;

    const span = end - start;
    if (span <= 0) return null;

    return { start, end, span, totalRows };
  };

  async function handleDownload() {
    const filename = selectedFile.name;

    // LVM
    if (selectedDownloadIndex === 0) {
      const link = document.createElement("a");
      link.href = `${endpointMapping.getLVMDownload}?file=${filename}.lvm`;
      link.setAttribute("download", filename);
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      return;
    }

    // Raw CSV
    if (selectedDownloadIndex === 1) {
      isLoadingPlotly = true;
      try {
        const a = document.createElement("a");
        a.href = `${endpointMapping.getExcelDownload}?file=${encodeURIComponent(
          filename
        )}.lvm`;
        a.setAttribute("download", `${filename}.xlsx`);
        a.click();
      } finally {
        isLoadingPlotly = false;
      }
      return;
    }

    // Filtered CSV (respecting current selection span)
    if (selectedDownloadIndex === 2) {
      const info = getSelectionInfo();
      if (!info) {
        plotError =
          "Please enter a valid Start Row and Row Count before exporting.";
        return;
      }
      const { start, span } = info;

      isLoadingPlotly = true;
      try {
        await exportFilteredExcel({
          selectedFile,
          xIndex: get(selectedXColumnIndex),
          yIndex: get(selectedYColumnIndex),
          startRow: start,
          numRows: span,
          filterNumber: selectedFilterIndex,
          sigma,
          windowSize,
          fetchStaticFireColumns,
          fetchStaticFireFilteredData
        });
      } catch (e) {
        console.error(e);
        plotError =
          e instanceof Error
            ? e.message
            : "Unexpected error while exporting filtered Excel.";
      } finally {
        isLoadingPlotly = false;
      }
    }
  }

  let data: Partial<Data>[] = [];

  const loadPlotly = async (datasets: Partial<Data>[]) => {
    const Plotly = await import("plotly.js-dist-min");
    await Plotly.newPlot(plotlyChartDiv, datasets, layout, config);
  };

  const onFullscreenChange = (isFullscreen: boolean) => {
    const bottomPadding = 80;

    if (isFullscreen) {
      layout.autosize = false;
      layout.width = window.innerWidth;
      layout.height = window.innerHeight - bottomPadding;
    } else {
      layout.autosize = true;
      layout.width = undefined;
      layout.height = shrunkenHeight;
    }

    loadPlotly(data);
  };

  const retrievefilterData = async (x: number[], y: number[]) => {
    if (selectedFilterIndex === 0) return null;

    const req: PostFilterDataRequest = {
      xColumns: x,
      yColumns: y,
      filterValue: selectedFilterIndex === 1 ? sigma : windowSize,
      filterNumber: selectedFilterIndex
    };

    const res = await fetchStaticFireFilteredData(req);
    if (!res) return null;

    return {
      x: res.xColumns,
      y: res.yColumns
    };
  };

  const fetchAndLoadPlotly = async (
    fetchData?: () => Promise<Partial<Data>[] | null>
  ) => {
    if (isLoadingPlotly) return;

    isLoadingPlotly = true;
    plotError = "";

    try {
      const newData = fetchData ? await fetchData() : [];
      data = newData ?? [];
      await loadPlotly(data);
    } catch (e) {
      console.error(e);
      plotError =
        e instanceof Error
          ? e.message
          : "Unexpected error while loading chart data.";
    } finally {
      isLoadingPlotly = false;
    }
  };

  // Parent hook – just redraw current data/layout
  $: refreshGraph = () => loadPlotly(data);

  const goPrev = () => {
    if (!isRollingMode) return;
    const info = getSelectionInfo();
    if (!info) return;

    pageOffset -= ROLLING_WINDOW;
    if (pageOffset < 0) pageOffset = 0;

    refreshPlotly();
  };

  const goNext = () => {
    if (!isRollingMode) return;
    const info = getSelectionInfo();
    if (!info) return;

    const { span } = info;
    pageOffset += ROLLING_WINDOW;

    const maxStartOffset = Math.max(span - ROLLING_WINDOW, 0);
    if (pageOffset > maxStartOffset) pageOffset = maxStartOffset;

    refreshPlotly();
  };

  const handleRefreshClick = () => {
    // Reset pagination when user changes selection
    pageOffset = 0;
    refreshPlotly();
  };

  export const refreshPlotly = async () => {
    if (!selectedFile) return;

    const info = getSelectionInfo();
    if (!info) {
      // Invalid selection: clear chart, show error, no fetch.
      isRollingMode = false;
      plotError = "Please enter a Row Count greater than Start Row.";
      data = [];
      await loadPlotly(data);
      return;
    }

    const { start, span } = info;

    // Pagination rule: if (Row Count - Start Row) > 250k, paginate.
    isRollingMode = span > ROLLING_WINDOW;

    // Clamp pageOffset inside the current selection span
    let effectiveOffset = pageOffset;
    if (!isRollingMode) {
      effectiveOffset = 0;
      pageOffset = 0;
    } else {
      const maxStartOffset = Math.max(span - ROLLING_WINDOW, 0);
      if (effectiveOffset < 0) effectiveOffset = 0;
      if (effectiveOffset > maxStartOffset) effectiveOffset = maxStartOffset;
      pageOffset = effectiveOffset;
    }

    const rowsThisPage = isRollingMode
      ? Math.min(ROLLING_WINDOW, span - effectiveOffset)
      : span;

    // Safety guard
    if (rowsThisPage <= 0) {
      isRollingMode = false;
      plotError = "Nothing to plot for this selection.";
      data = [];
      await loadPlotly(data);
      return;
    }

    const apiStartRow = start + effectiveOffset;

    const xColumnName =
      selectedFile.metadata.xColumnNames[get(selectedXColumnIndex)];
    const yColumnName =
      selectedFile.metadata.yColumnNames[get(selectedYColumnIndex)];

    const req: PostStaticFireColumnsRequest = {
      name: selectedFile.name,
      startRow: apiStartRow,
      numRows: rowsThisPage,
      xColumnNames: [xColumnName],
      yColumnNames: [yColumnName]
    };

    await fetchAndLoadPlotly(async () => {
      const res = await fetchStaticFireColumns(req);
      if (!res) return null;

      const xRaw = res.xColumns[xColumnName].rows as number[];
      const yRaw = res.yColumns[yColumnName].rows as number[];

      // Build shapes only if Test Start / End lie inside this page's x-range
      const numericX = xRaw
        .map((v) => Number(v))
        .filter((v) => Number.isFinite(v));

      let shapes: any[] = [];
      if (numericX.length > 0) {
        // *** FIX: avoid Math.min(...bigArray) / Math.max(...bigArray) ***
        let xMin = Infinity;
        let xMax = -Infinity;
        for (const v of numericX) {
          if (v < xMin) xMin = v;
          if (v > xMax) xMax = v;
        }

        const shapeDefs: { x: number; color: string; name: string }[] = [];

        if (testStart !== 0 && testStart >= xMin && testStart <= xMax) {
          shapeDefs.push({ x: testStart, color: "blue", name: "Test Start" });
        }
        if (testEnd !== 0 && testEnd >= xMin && testEnd <= xMax) {
          shapeDefs.push({ x: testEnd, color: "green", name: "Test End" });
        }

        shapes = shapeDefs.map((s) => ({
          type: "line",
          x0: s.x,
          x1: s.x,
          y0: 0,
          y1: 1,
          xref: "x",
          yref: "paper",
          name: s.name,
          showlegend: true,
          line: { color: s.color, width: 2, dash: "dash" }
        }));
      }
      layout.shapes = shapes;

      let datasets: Partial<Data>[] = [
        {
          x: xRaw,
          y: yRaw,
          type: "scattergl",
          mode: "lines",
          name: `${yColumnName} (Raw)`,
          line: { color: "#FFFFFF" }
        }
      ];

      if (selectedFilterIndex !== 0) {
        const filtered =
          (await retrievefilterData(xRaw, yRaw)) ?? { x: [], y: [] };
        const filterLabel =
          selectedFilterIndex === 1 ? "Gaussian" : "Moving Avg";

        const filteredDataset: Partial<Data> = {
          x: filtered.x,
          y: filtered.y,
          type: "scattergl",
          mode: "lines",
          name: `${yColumnName} (${filterLabel} Filter)`,
          line: { color: style.themeColor }
        };

        if (isFilterOn) {
          // Overlay filtered on top of raw
          datasets.push(filteredDataset);
        } else {
          // Only show filtered
          datasets[0] = filteredDataset;
        }
      }

      // Hover behaviour
      if (isRollingMode) {
        layout.hovermode = false;
        datasets.forEach((d) => {
          d.hoverinfo = "skip";
          d.hovertemplate = "";
        });
      } else {
        layout.hovermode = "closest";
        datasets.forEach((d) => {
          delete d.hoverinfo;
          delete d.hovertemplate;
        });
      }

      return datasets;
    });
  };

  // No auto plotting on mount – just mount an empty chart
  onMount(async () => {
    await loadPlotly([]);
  });
</script>

<div class="container">
  <div class="content-header">
    <div class="title">
      <h1>Dashboard for <i>{selectedFile.name}</i></h1>
      <p>
        Visualizing data for
        <i>{selectedFile.metadata.xColumnNames[$selectedXColumnIndex]}</i>
        and
        <i>{selectedFile.metadata.yColumnNames[$selectedYColumnIndex]}</i>
      </p>
    </div>
    <div class="data-select">
      <div class="column-select">
        <Dropdown
          onChange={(index) => ($selectedXColumnIndex = index)}
          isDisabled={isLoadingPlotly}
          label="X Column"
          id="x-column"
          options={selectedFile.metadata.xColumnNames}
        />
        <Dropdown
          onChange={(index) => ($selectedYColumnIndex = index)}
          isDisabled={isLoadingPlotly}
          label="Y Column"
          id="y-column"
          options={selectedFile.metadata.yColumnNames}
        />
        <Dropdown
          onChange={(index) => {
            selectedFilterIndex = index;
          }}
          isDisabled={isLoadingPlotly}
          label="Filter"
          id="filter"
          options={["None", "Gaussian", "Moving Average"]}
        />
      </div>

      <div class="filter-input">
        {#if selectedFilterIndex === 1}
          <Input
            id="sigma"
            placeholder="0"
            isDisabled={isLoadingPlotly}
            label="Sigma"
            onChange={(value) => (sigma = safeParseInt(value))}
          />
        {/if}
        {#if selectedFilterIndex === 2}
          <Input
            id="windowSize"
            placeholder="0"
            isDisabled={isLoadingPlotly}
            label="Window Size"
            onChange={(value) => (windowSize = safeParseInt(value))}
          />
        {/if}
      </div>

      <div class="time-select">
        <Input
          id="test-start"
          placeholder="0"
          isDisabled={isLoadingPlotly}
          label="Test Start"
          onChange={(value) => (testStart = safeParseFloat(value))}
        />
        <Input
          id="test-end"
          placeholder="0"
          isDisabled={isLoadingPlotly}
          label="Test End"
          onChange={(value) => (testEnd = safeParseFloat(value))}
        />
      </div>

      <div class="row-select">
        <Input
          id="start-row"
          placeholder="0"
          isDisabled={isLoadingPlotly}
          label="Start Row"
          regex={numericRegex}
          onChange={(value) => {
            startRow = safeParseInt(value);
            pageOffset = 0;
          }}
        />
        <Input
          id="num-rows"
          placeholder={selectedFile.metadata.totalRows.toString()}
          isDisabled={isLoadingPlotly}
          label="Row Count"
          regex={numericRegex}
          onChange={(value) => {
            numRows = safeParseInt(value);
            pageOffset = 0;
          }}
        />
      </div>
    </div>
  </div>

  <div class="content-container">
    <div class="chart-pod pod" bind:this={fullscreenTarget}>
      <div class="title-container">
        <h2>Static Fire Chart</h2>
        <div style="display: flex; gap: 0.5rem;">
          <div>
            <DropdownBtn
              onChange={(index) => {
                selectedDownloadIndex = index;
              }}
              Click={handleDownload}
              label="Download"
              isDisabled={isLoadingPlotly}
              id="download"
              buttonize={true}
              icon={Download}
              options={["LVM", "Raw CSV", "Filtered CSV"]}
            />
          </div>

          {#if isRollingMode}
            <IconButton icon={ChevronLeft} onClick={goPrev} />
            <IconButton icon={ChevronRight} onClick={goNext} />
          {/if}

          <IconButton icon={ChartLine} onClick={handleFilter} toggle={true} />
          <IconButton icon={RefreshCcw} onClick={handleRefreshClick} />
          <FullscreenButton
            onChange={onFullscreenChange}
            targetElement={fullscreenTarget}
          />
        </div>
      </div>
      <div class="chart-wrapper">
        <div
          class="loading-overlay"
          class:hidden={!isLoadingPlotly && !plotError}
        >
          {#if plotError}
            <div>
              <MessageCircleWarningIcon />
              <b>{plotError}</b>
            </div>
          {:else}
            <Loader2 />
          {/if}
        </div>
        <div bind:this={plotlyChartDiv} class="chart"></div>
      </div>
    </div>
    <div class="value-pods">
      <div class="min-val-pod pod">
        <label for="min-val">Minimum Value</label>
        <div class="value" id="min-val">0.00</div>
      </div>
      <div class="max-val-pod pod">
        <label for="max-val">Maximum Value</label>
        <div class="value" id="max-val">0.00</div>
      </div>
      <div class="avg-val-pod pod">
        <label for="avg-val">Average Value</label>
        <div class="value" id="avg-val">0.00</div>
      </div>
    </div>
  </div>
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  .toggle-container.active {
    background-color: $bg-color-highlighted;
  }

  .container {
    flex-grow: 1;
    padding: 1rem;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
  }

  div.content-header {
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    gap: 2rem;

    div.title {
      margin-right: auto;
    }

    div.container:hover > button.dropdown-button {
      background-color: $bg-color-highlighted;
      border-color: $txt-color-highlighted;

      :global(.lucide-icon) {
        stroke: $txt-color-highlighted;
      }

      span.label {
        color: $txt-color-highlighted;
      }
    }

    div.data-select {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      width: 100%;
      gap: 1rem;
      margin-bottom: 0.8rem;

      & > div {
        gap: 1.5rem;
        flex-grow: 1;
        display: flex;
        flex-direction: row;
      }

      & > div.row-select {
        justify-content: flex-end;
      }
    }
  }

  div.pod {
    background-color: $bg-color-4;
    border: 1px solid $outline-color-1;
    border-radius: $border-radius-1;
    padding: 1rem;

    h2 {
      margin: 0;
      margin-bottom: 0.5em;
    }

    div.value {
      margin-top: 0.5rem;
      font-size: 1.5rem;
      font-weight: bold;
    }
  }

  div.content-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding-bottom: 1rem;

    div.value-pods {
      display: flex;
      flex-direction: row;
      gap: 1rem;

      div.pod {
        flex-grow: 1;
      }
    }

    div.chart-pod {
      flex-grow: 1;

      div.title-container {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 0.7rem;

        h2 {
          margin: 0;
        }
      }

      div.chart-wrapper {
        border-radius: $border-radius-1;
        overflow: hidden;
        position: relative;

        div.loading-overlay {
          position: absolute;
          top: 0;
          left: 0;
          width: 100%;
          height: 100%;
          display: flex;
          justify-content: center;
          align-items: center;
          background-color: rgba(0, 0, 0, 0.5);
          z-index: 1;

          :global(.lucide-icon) {
            width: 3.3rem;
            height: auto;
          }

          &.hidden {
            display: none;
          }

          & > div {
            display: flex;
            flex-direction: column;
            gap: 0.6rem;
            justify-content: center;
            align-items: center;

            b {
              color: $txt-color-highlighted;
              font-size: 1.2rem;
            }

            :global(.lucide-icon) {
              width: 3rem;
              height: auto;
              stroke: $txt-color-highlighted;
            }
          }
        }
      }
    }
  }
</style>
