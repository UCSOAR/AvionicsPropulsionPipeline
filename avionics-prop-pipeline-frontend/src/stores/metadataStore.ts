import { defineStore } from 'pinia'

export const useMetadataStore = defineStore('metadata', {
  state: () => ({
    metadata: {} as Record<string, any>
  }),
  actions: {
    setMetadata(newMetadata: Record<string, any>) {
      this.metadata = newMetadata
    }
  }
})
