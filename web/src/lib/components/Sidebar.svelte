<script lang="ts">
  import UploadFile from "$lib/components/UploadFile.svelte";
  import IconButton from "./IconButton.svelte";
  import FolderTree, { type Node } from "$lib/components/FolderTree.svelte";
  import { onMount, tick } from "svelte";
  import { PanelLeftClose, PanelLeftOpen, FolderPlus, UploadCloud } from "@lucide/svelte";
  import { endpointMapping } from "$lib/utils/constants";
  import type { SelectedFile } from "$lib/models/selectedFile";
  import CreateFileModal from "$lib/components/CreateFileModal.svelte";
  import UploadFileModal from "./UploadFileModal.svelte";


  export let selectedFile: SelectedFile | undefined = undefined;
  export let refreshDashboardGraph: () => Promise<void>;
  export let isExpanded = true;

  let files: Record<string, any> = {};
  let tree: Node[] = [];
  let error: string | null = null;
  let showCreateModal = false;
  let showUploadModal = false;
  /** Toggle sidebar open/close */
  const toggleSidebar = () => {
    isExpanded = !isExpanded;
    refreshDashboardGraph();
  };

  /** When a file is selected */
const handleFileSelect = async (path: string, metadata: any, type: string) => {
  if (type !== "file") {
    // clicked a folder → don’t trigger dashboard update
    console.log("📁 Folder selected:", path);
    return;
  }

  selectedFile = { name: path, metadata };
  await tick(); // ensure DOM ready
  refreshDashboardGraph?.();
};


  /** Converts flat path map → nested folder structure */
function buildTree(map: Record<string, any>): Node[] {
  const root: Node[] = [];

  function insert(path: string, metadata: any) {
    // Normalize slashes and remove trailing ones
    const cleanPath = path.replace(/\\/g, "/").replace(/\/+$/, "");
    const parts = cleanPath.split("/").filter(Boolean);

    let current = root;
    let accumulated = "";

    for (let i = 0; i < parts.length; i++) {
      const part = parts[i];
      accumulated = accumulated ? `${accumulated}/${part}` : part;

      let existing = current.find((n) => n.name === part);

      const isLastPart = i === parts.length - 1;
      const isFolder = metadata === null || !isLastPart;

      // If this node doesn't exist yet, create it
      if (!existing) {
        existing = {
          name: part,
          path: accumulated,
          type: isFolder ? "folder" : "file",
          children: isFolder ? [] : undefined,
        };
        current.push(existing);
      }

      // If it's a folder, go deeper into its children
      if (isFolder) {
        current = existing.children!;
      } else {
        // Attach metadata only to files
        existing.metadata = metadata;
      }
    }
  }

  // Iterate through every entry in the map
  Object.entries(map).forEach(([path, meta]) => insert(path, meta));

  return root;
}


  /** Fetch files and build tree */
  const fetchFiles = async () => {
    try {
      const response = await fetch(endpointMapping.getStaticFireMetadataUrl, {
        method: "GET",
        credentials: "include",
        headers: { "Content-Type": "application/json" },
      });

      if (!response.ok) throw new Error("Failed to fetch files");

      files = await response.json();
      console.log(files);
      tree = buildTree(files);
      error = null;
    } catch (err) {
      error = err instanceof Error ? err.message : "An error occurred.";
      tree = [];
    }
  };

  /** Create a new folder in the backend */
  const createFolder = async (path: string, name: string) => {
    try {
      const response = await fetch(endpointMapping.createFileOrFolderUrl, {
        method: "POST",
        credentials: "include",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          type: "folder", // ✅ required field for backend switch
          path,
          name,
        }),
      });

      if (!response.ok) {
        throw new Error(`Failed to create folder: ${response.statusText}`);
      }

      await fetchFiles();
    } catch (err) {
      error = err instanceof Error ? err.message : "An error occurred.";
    }
  };



  onMount(fetchFiles);

  /** Refresh after upload */
  const handleUploadComplete = () => {
    fetchFiles();
  };
</script>

<aside class="side-bar {isExpanded ? 'expanded' : 'collapsed'}" on:transitionend={refreshDashboardGraph}>
  <!-- Upload Section -->
  <div class="upload-container">
  <IconButton
  icon={UploadCloud}
  onClick={() => (showUploadModal = true)}
/>

{#if showUploadModal}
  <UploadFileModal
    tree={tree}
    onClose={() => (showUploadModal = false)}
    onUploadComplete={fetchFiles}
/>
{/if}
  <IconButton
    icon={FolderPlus}
    onClick={() => (showCreateModal = true)}
  />
</div>


  <!-- Header -->
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

  <!-- File Tree -->
  <div class="file-list">
    {#if tree.length > 0}
      <FolderTree
        nodes={tree}
        selected={selectedFile?.name ?? ""}
        on:select={(e) => handleFileSelect(e.detail.path, e.detail.metadata, e.detail.type)}
        isCollapsed={!isExpanded}
      />
    {:else if isExpanded}
      <p class="empty">{error || "No uploaded files yet."}</p>
    {/if}
  </div>

  {#if showCreateModal}
    <CreateFileModal
      tree={tree}
      onClose={() => (showCreateModal = false)}
      onCreate={async (path, name) => {
        await createFolder(path, name);
        await fetchFiles();
        showCreateModal = false;
      }}
    />
  {/if}

</aside>

<style lang="scss">
  @use "../styles/variables.scss" as *;

  aside.side-bar {
    display: flex;
    flex-direction: column;
    background-color: #121212;
    height: 100vh;
    overflow: hidden;
    border-right: 1px solid $outline-color-1;
    transition: all 0.2s ease;

    &.expanded {
      min-width: 20rem;
    }

    &.collapsed {
      min-width: 4.5rem;

      .upload-container {
        display: none;
      }
    }
  }

  .upload-container {
    padding: 1rem;
    white-space: nowrap;
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

  .empty {
    color: #aaa;
    font-size: 0.9rem;
    padding: 1rem;
  }
</style>
