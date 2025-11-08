<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import FolderTree, { type Node } from "$lib/components/FolderTree.svelte";
  import IconButton from "$lib/components/IconButton.svelte";
  import { UploadCloud, X } from "@lucide/svelte";
  import { endpointMapping } from "$lib/utils/constants";

  export let tree: Node[] = [];
  export let onClose: () => void;
  export let onUploadComplete: () => void;

  const dispatch = createEventDispatcher();

  let selectedPath = "";
  let fileInput: HTMLInputElement;
  let uploading = false;
  let selectedFile: File | null = null;

  const handleFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (!target.files?.length) return;
    selectedFile = target.files[0];
  };

  const handleUpload = async () => {
    if (!selectedFile) return alert("Please choose a .lvm file first!");

    uploading = true;

    const formData = new FormData();
    formData.append("file", selectedFile);
    formData.append("path", selectedPath);

    try {
      const res = await fetch(endpointMapping.uploadStaticFireUrl, {
        method: "POST",
        credentials: "include",
        body: formData,
      });

      if (!res.ok) throw new Error(await res.text());
      console.log(`✅ Uploaded ${selectedFile.name} to ${selectedPath || "(root)"}`);
      onUploadComplete();
      onClose();
    } catch (err) {
      console.error("Upload failed:", err);
      alert("Upload failed: " + err);
    } finally {
      uploading = false;
    }
  };
</script>

<div class="overlay" on:click={(e) => e.target === e.currentTarget && onClose()}>
  <div class="modal" on:click|stopPropagation>
    <div class="modal-header">
      <h2>Upload .lvm File</h2>
      <IconButton icon={X} onClick={onClose} />
    </div>

    <div class="modal-body">
      <label>Select Destination Folder:</label>
      <div class="folder-tree">
        <FolderTree
          nodes={tree.filter((n) => n.type === "folder")}
          selected={selectedPath}
          on:select={(e) => (selectedPath = e.detail.path)}
          isCollapsed={false}
        />
      </div>

      <label>Choose File:</label>
      <input
        type="file"
        accept=".lvm"
        bind:this={fileInput}
        on:change={handleFileChange}
      />

      {#if selectedPath}
        <p class="selected-path">📁 Uploading to: <b>{selectedPath}</b></p>
      {:else}
        <p class="selected-path muted">📁 Uploading to root directory</p>
      {/if}

      <div class="modal-footer">
        <IconButton
          icon={UploadCloud}
          label={uploading ? "Uploading..." : "Upload"}
          onClick={handleUpload}
          isDisabled={uploading}
        />
        <button class="cancel" on:click={onClose}>Cancel</button>
      </div>
    </div>
  </div>
</div>

<style lang="scss">
  .overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.65);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 50;
  }

  .modal {
    background: #1c1c1c;
    border-radius: 8px;
    width: 420px;
    max-height: 80vh;
    overflow-y: auto;
    padding: 1rem;
    color: white;
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .folder-tree {
    max-height: 220px;
    overflow-y: auto;
    border: 1px solid rgba(255,255,255,0.1);
    border-radius: 4px;
    padding: 0.3rem;
    margin-bottom: 1rem;
  }

  input[type="file"] {
    margin-top: 0.5rem;
    margin-bottom: 0.8rem;
  }

  .selected-path {
    font-size: 0.85rem;
    color: #ccc;
    margin-bottom: 0.8rem;
    &.muted { opacity: 0.7; }
  }

  .modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 0.5rem;
    .cancel {
      background: #333;
      color: #ccc;
      border: none;
      border-radius: 4px;
      padding: 0.5rem 1rem;
      cursor: pointer;
    }
  }
</style>
