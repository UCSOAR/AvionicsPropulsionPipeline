<script lang="ts">
  import "$lib/styles/global.scss";
  import "$lib/components/IconButton.svelte";
  import type { SelectedFile } from "$lib/models/selectedFile";
  import Dashboard from "$lib/components/Dashboard.svelte";
  import TopBar from "$lib/components/TopBar.svelte";
  import Sidebar from "$lib/components/Sidebar.svelte";

  let selectedFile: SelectedFile | undefined = undefined;
</script>

<main class="app-container">
  <TopBar />

  <div class="main-layout">
    <Sidebar bind:selectedFile />

    {#if selectedFile}
      <Dashboard {selectedFile} />
    {/if}
  </div>
</main>

<style lang="scss">
  @use "../../lib/styles/variables.scss" as *;

  main.app-container {
    display: flex;
    flex-direction: column;
    height: 100vh; /* fill the entire viewport */
    width: 100%;
  }

  .main-layout {
    display: flex;
    flex: 1; /* take the remaining height after TopBar */
    overflow: hidden; /* avoid scrollbars unless necessary */
  }

  /* Optional: make sure Sidebar and Dashboard stretch full height */
  :global(.side-bar) {
    height: 100%;
  }

  :global(.dashboard-container) {
    flex: 1;
    overflow-y: auto;
  }
</style>