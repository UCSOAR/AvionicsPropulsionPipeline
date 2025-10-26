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


export type PostFilterDataRequest = {
  xColumns: number[];
  yColumns: number[];
  filterValue: number;
  filterNumber: number;
};

export type PostFilterDataResponse = {
  xColumns: number[];
  yColumns: number[];
};


