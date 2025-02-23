interface TimestampMetadata {
  date: string;
  time: string;
}

export interface PreviewMetadata {
  resultTimestamp: TimestampMetadata;
  operator: string;
  xColumnNames: string[];
  yColumnNames: string[];
}
