<script lang="ts">
  import type { PreviewMetadata } from "$lib/models/cacheTreeModels";
  import Dropdown from "./Dropdown.svelte";

  type SelectedFile = {
    name: string;
    metadata: PreviewMetadata;
  };

  export let selectedFile: SelectedFile | null = null;
</script>

<div class="container">
  {#if selectedFile}
    <div class="content-header">
      <div class="title">
        <h1>Dashboard for <i>{selectedFile.name}</i></h1>
        <p>
          Visualizing data for <i>{selectedFile.metadata.xColumnNames[0]}</i>
          and
          <i>{selectedFile.metadata.yColumnNames[0]}</i>
        </p>
      </div>
      <div class="column-select">
        <Dropdown
          label="X Column"
          id="x-column"
          options={selectedFile.metadata.xColumnNames}
        />
        <Dropdown
          label="Y Column"
          id="y-column"
          options={selectedFile.metadata.yColumnNames}
        />
      </div>
    </div>
  {:else}
    <h1>Dashboard</h1>
    <div class="message-container">
      <p>Select a file from the sidebar to view its data.</p>
    </div>
  {/if}
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.content-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-direction: row;
    gap: 1rem;

    div.column-select {
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
      margin-right: 1rem;
    }
  }

  div.container {
    padding: 1rem;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    flex-grow: 1;

    h1 {
      margin: 0;
      margin-bottom: 0.3rem;

      i {
        color: $txt-color-1;
      }
    }

    div.message-container {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-grow: 1;
      text-align: center;
    }
  }
</style>
