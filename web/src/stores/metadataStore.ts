import { writable } from 'svelte/store';

interface Metadata {
  xColumnNames: string[];
  yColumnNames: string[];
  [key: string]: any;
}

export const metadataStore = writable<{
  metadata: Metadata | null;
  xColumn: string | null;
  yColumn: string | null;
  fileName: string | null;
}>({
  metadata: null,
  xColumn: null,
  yColumn: null,
  fileName: null
});

export function setMetadata(metadata: Metadata, xCol: string, yCol: string, fileName: string) {
  metadataStore.set({
    metadata,
    xColumn: xCol,
    yColumn: yCol,
    fileName
  });
}
