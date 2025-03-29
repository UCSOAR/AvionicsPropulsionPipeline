<script lang="ts">
<<<<<<< HEAD
  import { onMount} from 'svelte';
  import { ArrowUpToLine, PanelLeftClose, PanelLeftOpen, File } from 'lucide-svelte';
  import { useMetadataStore } from '../../stores/metadataStore';
  import FileUploader from '$lib/components/UploadFile.svelte';
  import { endpointMapping } from "$lib/utils/constants";

  export let data = {};
=======
  import FileUploader from "$lib/components/UploadFile.svelte";
  import type { SelectedFile } from "$lib/models/selectedFile";
  import { onMount } from "svelte";
  import { PanelLeftClose, PanelLeftOpen, File } from "@lucide/svelte";
  import { endpointMapping } from "$lib/utils/constants";

  export let selectedFile: SelectedFile | undefined = undefined;
>>>>>>> c0f757bf2cf4000fa6d7e28c302761699b53fb7b

  let isExpanded = true;
  let files: Record<string, any> = {};
  let error: string | null = null;
<<<<<<< HEAD
  let xCol: string | null = null;
  let yCol: string | null = null;
  let metadata: Record<string, any> = {};
  let fileName: string | null = null;

=======
>>>>>>> c0f757bf2cf4000fa6d7e28c302761699b53fb7b

  const toggleSidebar = () => {
    isExpanded = !isExpanded;
  };

  const handleFileClick = (fileName: string, metadata: any) => {
<<<<<<< HEAD
    console.log(`File clicked: ${fileName}`, metadata);
    data.fileName = fileName;
    data.metadata = metadata;

    

=======
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
        },
      };
    }

    selectedFile.name = fileName;
    selectedFile.metadata = metadata;
>>>>>>> c0f757bf2cf4000fa6d7e28c302761699b53fb7b
  };

  const fetchFiles = async () => {
    try {
      const response = await fetch(endpointMapping.getStaticFireMetadataUrl);
      if (!response.ok) throw new Error("Failed to fetch files");
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
<<<<<<< HEAD
    // Refetch the files after upload completes
=======
>>>>>>> c0f757bf2cf4000fa6d7e28c302761699b53fb7b
    fetchFiles();
  };
</script>

<aside class="side-bar {isExpanded ? 'expanded' : 'collapsed'}">
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
<<<<<<< HEAD
        <button 
          class="file-item" 
=======
        <button
          class="file-item"
>>>>>>> c0f757bf2cf4000fa6d7e28c302761699b53fb7b
          on:click={() => handleFileClick(name, metadata)}
        >
          <File size={16} class="icon" />
          {#if isExpanded}
            <span class="file-name">{name.replace(/\.[^.]+$/, "")}</span>
          {/if}
        </button>
      {/each}
    {:else}
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
<<<<<<< HEAD

=======
>>>>>>> c0f757bf2cf4000fa6d7e28c302761699b53fb7b
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
<<<<<<< HEAD
</style>
=======
</style>
>>>>>>> c0f757bf2cf4000fa6d7e28c302761699b53fb7b
