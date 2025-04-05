<script lang="ts">
  import FileUploader from "$lib/components/UploadFile.svelte";
  import type { SelectedFile } from "$lib/models/selectedFile";
  import { onMount } from "svelte"; 
  import { PanelLeftClose, PanelLeftOpen, File,  RefreshCcw} from "@lucide/svelte";
  import { endpointMapping } from "$lib/utils/constants";
  export let selectedFile: SelectedFile | undefined = undefined;
  export let refreshDashboardGraph: () => Promise<void>;


  export let isExpanded = true;
  let files: Record<string, any> = {};
  let error: string | null = null;

  const toggleSidebar = () => {
    isExpanded = !isExpanded;
  };

  const handleFileClick = (fileName: string, metadata: any) => {
    if (!selectedFile) {
      selectedFile = {
        name: "",
        metadata: {
          operator: "",
          resultTimestamp: {
            date: "",
            time: "",
          },
          xColumnNames: [],
          yColumnNames: [],
          totalRows: 0,
        },
      };
    }

    selectedFile.name = fileName;
    selectedFile.metadata = metadata;
  };

  const fetchFiles = async () => {
    try {
      const response = await fetch(endpointMapping.getStaticFireMetadataUrl);
      if (!response.ok && isExpanded  ) throw new Error("Failed to fetch files");
      files = await response.json();
      error = null;
    } catch (err) {
      error = err instanceof Error ? err.message : "An error occurred.";
    }
  };

  onMount(() => {
    fetchFiles();
  });

  const handleUploadComplete = () => {
    fetchFiles();
  };


</script>

<aside class="side-bar {isExpanded ? 'expanded' : 'collapsed'}" on:transitionend={refreshDashboardGraph}>
  <!-- Upload Section -->
  <div class="upload-container">
    <FileUploader onUploadComplete={handleUploadComplete} />
  </div>

  <!-- Files Header -->
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

  <!-- File List -->
  <div class="file-list">
    {#if Object.keys(files).length > 0}
      {#each Object.entries(files) as [name, metadata]}
        <button
          class={`${isExpanded ?"file-item" : "icon-item"}  ${selectedFile?.name === name ? "selected" : ""}`}
          on:click={() => handleFileClick(name, metadata)}
        >
          <File size={16} color={selectedFile?.name === name ? "#e64d4d" : "white"}/>
          {#if isExpanded}
            <span class="file-name">{name.replace(/\.[^.]+$/, "")}</span>
          {/if}  
        </button>
      {/each}
    {:else}

    {#if isExpanded}
      <p class="empty">{error || "No uploaded files yet."}</p>
    {/if}
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
    border-right: 1px solid $outline-color-1;

    &.expanded {
      width: 280px;
    }

    &.collapsed {
      width: 60px;
    }
  }

  .upload-container {
    width: 16rem;
    padding: 1rem;
    display: flex;
    justify-content: center;
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

    .file-name {
      font-size: 0.9rem;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }

    .side-bar.expanded .file-name {
      opacity: 1;
      transform: translateX(0);
    }

    &:hover {
      background-color: $bg-color-highlighted;
    }

    &:hover .file-name{
      color: $txt-color-highlighted
    }
    
  }

  .icon-item {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    color: white;
    padding: 0.5rem;
    border-radius: 4px;
    border: none;
    background: transparent;
    cursor: pointer;

    &:hover {
      background-color: $bg-color-highlighted;
    }

    &:hover .icon{
      color: $txt-color-highlighted
    }

  }

  .empty {
    color: #aaa;
    font-size: 0.9rem;
    padding: 1rem;
  }

  .selected {
  background-color: $bg-color-highlighted;
  color: $txt-color-highlighted;

  .file-name,
  .icon {
    color: $txt-color-highlighted;
  }
}


</style>

