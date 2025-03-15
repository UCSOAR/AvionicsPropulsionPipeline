export type TimestampMetadata = {
  data: string;
  time: string;
};

export type PreviewMetadata = {
  resultTimestamp: TimestampMetadata;
  operator: string;
  xColumnNames: string[];
  yColumnNames: string[];
};

export type ColumnNode = {
  rows: number[];
};

export type YColumnMetadata = {
  samples: number;
  date: string;
  unitLabel: string;
  xDimension: string;
};
