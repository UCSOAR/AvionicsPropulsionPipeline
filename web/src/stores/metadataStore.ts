import { defineStore } from 'pinia';

export const useMetadataStore = defineStore('metadata', {
  state: () => ({
    metadata: {} as Record<string, any>, // Holds metadata object
    colX: null as string | null,        // Selected column for X
    colY: null as string | null,        // Selected column for Y
    name: null as string | null         // The name of the file selected
  }),
  actions: {
    /**
     * Sets the metadata and the selected columns for X and Y.
     * @param newMetadata - The new metadata object to store.
     * @param X - The column name to set for the X-axis.
     * @param Y - The column name to set for the Y-axis.
     */
    setMetadata(newMetadata: Record<string, any>, X: string, Y: string, name: string) {
      this.metadata = newMetadata; // Update metadata
      this.colX = X;               // Update selected X column
      this.colY = Y;               // Update selected Y column
      this.name = name;
    },
  },
});
