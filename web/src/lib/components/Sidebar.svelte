<script lang="ts">
  import { ArrowUpToLine, PanelLeftClose, PanelLeftOpen } from 'lucide-svelte';
  import { onMount } from 'svelte';

  let isExpanded = true;
  let files: Array<{ name: string }> = [];

  const toggleSidebar = () => {
    isExpanded = !isExpanded;
  };

  const fetchFiles = async () => {
    files = [
      { name: 'RHT_2023-07-14-12-40_PM_shortened' },
      { name: 'RHT_2023-07-15-09-10_PM_shortened' }
    ];
  };

  const handleFileClick = (fileName: string) => {
    console.log(`File clicked: ${fileName}`);
  };

  onMount(() => {
    fetchFiles();
  });
</script>

<aside class="side-bar {isExpanded ? 'expanded' : 'collapsed'}">
  <div class="upload-container">
    {#if isExpanded}
      <button class="upload-button" on:click={() => console.log('Upload file')}>
        <ArrowUpToLine size={20} />
        <span class="label">Upload File</span>
      </button>
    {:else}
      <button class="upload-button icon-only" on:click={() => console.log('Upload file')}>
        <ArrowUpToLine size={20} />
      </button>
    {/if}
  </div>

  <div class="files-header">
    {#if isExpanded}
      <h3>Files</h3>
    {/if}
    <button class="layout-button" on:click={toggleSidebar}>
      {#if isExpanded}
        <PanelLeftClose size={20} />
      {:else}
        <PanelLeftOpen size={20} />
      {/if}
    </button>
  </div>

  <div class="file-list">
    {#if files.length > 0}
      {#each files as file}
        <button class="file-item" on:click={() => handleFileClick(file.name)}>
          {#if isExpanded}
            <ArrowUpToLine size={16} class="icon" />
            <span class="file-name">{file.name}</span>
          {:else}
            <ArrowUpToLine size={16} class="icon" />
          {/if}
        </button>
      {/each}
    {:else}
      <p class="empty">No uploaded files yet.</p>
    {/if}
  </div>
</aside>

<style lang="scss">
  @use "../styles/variables.scss" as *;

  .side-bar {
    display: flex;
    flex-direction: column;
    background-color: #121212;
    box-shadow: 2px 0 10px rgba(0, 0, 0, 0.3);
    transition: width 0.3s ease;
    height: 100vh;
    overflow: hidden;

    &.expanded {
      width: 280px;
    }

    &.collapsed {
      width: 60px;
    }
  }

  .upload-container {
    padding: 1rem;
    display: flex;
    justify-content: center;
  }

  .upload-button {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #333;
    border-radius: 8px;
    background-color: transparent;
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    cursor: pointer;

    &.icon-only {
      width: 40px;
      height: 35px;
      padding: 0;
      border-radius: 5px;
    }
  }

  .files-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.75rem 1rem;

    h3 {
      font-size: 1.25rem;
      font-weight: 600;
      color: white;
    }

    .layout-button {
      background: transparent;
      border: none;
      color: white;
      padding: 0.25rem;
      cursor: pointer;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }

  .file-list {
    flex: 1;
    overflow-y: auto;
    padding: 0 0.5rem;
    display: flex;
    flex-direction: column;
  }

  .file-item {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: white;
    padding: 0.5rem;
    border-radius: 4px;
    border: none;
    background: transparent;
    cursor: pointer;

    &:hover {
      background-color: rgba(255, 255, 255, 0.1);
    }

    .file-name {
      font-size: 0.9rem;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }

  .empty {
    color: #aaa;
    font-size: 0.9rem;
    padding: 1rem;
  }
</style>
