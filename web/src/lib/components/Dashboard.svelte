<script lang="ts">
  import type { PreviewMetadata } from "$lib/models/cacheTreeModels";

  export let metadata: PreviewMetadata | null = null;
</script>

<div class="container">
  <h1>Dashboard</h1>
  {#if metadata}
    <pre>{JSON.stringify(metadata, null, 2)}</pre>
  {:else}
    <div class="message-container">
      <p>Select a file from the sidebar to view its data.</p>
    </div>
  {/if}
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.container {
    padding: 1rem;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    flex-grow: 1;

    h1 {
      margin: 0;
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
