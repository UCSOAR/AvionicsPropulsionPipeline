<script lang="ts">
  import UploadFile from "$lib/components/UploadFile.svelte";
  import IconButton from "./IconButton.svelte";
  import { onMount } from "svelte";
  import { PanelLeftClose, PanelLeftOpen, File } from "@lucide/svelte";
  import { endpointMapping } from "$lib/utils/constants";
  import type { SelectedFile } from "$lib/models/selectedFile";

  export let selectedFile: SelectedFile | undefined = undefined;
  export let refreshDashboardGraph: () => Promise<void>;

  export let isExpanded = true;
  let files: Record<string, any> = {};
  let error: string | null = null;

  const toggleSidebar = () => {
    isExpanded = !isExpanded;
    refreshDashboardGraph();
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
      if (!response.ok && isExpanded) throw new Error("Failed to fetch files");
      files = await response.json();
      error = null;
    } catch (err) {
      error = err instanceof Error ? err.message : "An error occurred.";
    }
  };

  onMount(fetchFiles);

  const handleUploadComplete = () => {
    fetchFiles();
  };
</script>

<aside class="side-bar {isExpanded ? 'expanded' : 'collapsed'}">
  <!-- Upload Section -->
  <div class="upload-container">
    <UploadFile onUploadComplete={handleUploadComplete} />
  </div>

  <!-- Files Header -->
  <div class="files-header">
    {#if isExpanded}
      <h3>Files</h3>
    {/if}
    <div class="button-container">
      {#if isExpanded}
        <IconButton icon={PanelLeftClose} onClick={toggleSidebar} />
      {:else}
        <IconButton icon={PanelLeftOpen} onClick={toggleSidebar} />
      {/if}
    </div>
  </div>

  <!-- File List -->
  <div class="file-list">
    {#if Object.keys(files).length > 0}
      {#each Object.entries(files) as [name, metadata]}
        <button
          class={`${isExpanded ? "file-item" : "icon-item"}  ${selectedFile?.name === name ? "selected" : ""}`}
          on:click={() => handleFileClick(name, metadata)}
        >
          <File
            size={16}
            color={selectedFile?.name === name ? "#e64d4d" : "white"}
          />
          {#if isExpanded}
            <span class="file-name">{name.replace(/\.[^.]+$/, "")}</span>
          {/if}
        </button>
      {/each}
    {:else if isExpanded}
      <p class="empty">{error || "No uploaded files yet."}</p>
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
    height: 100vh;
    overflow: hidden;
    border-right: 1px solid $outline-color-1;

    &.expanded {
      width: 20rem;
    }

    &.collapsed {
      width: 4.5rem;
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
  }

  .file-list {
    flex: 1;
    overflow-y: auto;
    padding: 0 0.5rem;
    gap: 0.3rem;
    display: flex;
    flex-direction: column;
  }

  .icon-item {
    display: flex;
    align-items: center;
    justify-content: center;
    color: $txt-color-1;
    padding: 0.5rem;
    border-radius: $border-radius-1;
    border: none;
    background: transparent;
    cursor: pointer;

    &:hover {
      background-color: $bg-color-highlighted;
    }
  }

  .empty {
    color: #aaa;
    font-size: 0.9rem;
    padding: 1rem;
  }

  .file-item {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    color: $txt-color-1;
    padding: 0.5rem;
    border-radius: $border-radius-1;
    border: none;
    background: transparent;
    cursor: pointer;

    .file-name {
      font-size: 0.9rem;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }

    &:hover {
      background-color: $bg-color-highlighted;

      .file-name {
        color: $txt-color-highlighted;
      }

      :global(.lucide-icon) {
        stroke: $txt-color-highlighted;
      }
    }
  }

  .selected {
    background-color: $bg-color-highlighted;
    color: $txt-color-highlighted;

    .file-name {
      color: $txt-color-highlighted;
    }
  }
</style>
