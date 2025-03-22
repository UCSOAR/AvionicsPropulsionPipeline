import type { ColumnNode, YColumnMetadata } from "./cacheTreeModels";

export type PostStaticFireColumnsRequest = {
  name: string;
  startRow: number;
  numRows: number;
  xColumnNames: string[];
  yColumnNames: string[];
};

export type PostStaticFireColumnsResponse = {
  yColumnMetadata: Record<string, YColumnMetadata>;
  xColumns: Record<string, ColumnNode>;
  yColumns: Record<string, ColumnNode>;
};
