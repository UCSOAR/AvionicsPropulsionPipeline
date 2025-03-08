import { create } from "zustand";

// Define the metadata store
interface MetadataState {
  metadata: Record<string, any>; // Holds metadata object
  colX: string | null;           // Selected column for X
  colY: string | null;           // Selected column for Y
  name: string | null;           // Selected file name
  setMetadata: (newMetadata: Record<string, any>, X: string, Y: string, name: string) => void;
}

// Create Zustand store
export const useMetadataStore = create<MetadataState>((set) => ({
  metadata: {},    // Default empty metadata
  colX: null,      // Default null X column
  colY: null,      // Default null Y column
  name: null,      // Default null file name

  // Set metadata function (equivalent to Pinia action)
  setMetadata: (newMetadata, X, Y, name) => {
    set({ metadata: newMetadata, colX: X, colY: Y, name });
  },
}));
