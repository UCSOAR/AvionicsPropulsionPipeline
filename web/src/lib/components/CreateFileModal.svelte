<script lang="ts">
  import FolderTree, { type Node } from "$lib/components/FolderTree.svelte";
  import IconButton from "$lib/components/IconButton.svelte";
  import { X } from "@lucide/svelte";

  export let tree: Node[] = [];
  export let onClose: () => void;
  export let onCreate: (path: string, name: string) => void;

  let selectedPath = "";
  let newName = "";

  /** Filter out files so we only display folders */
  function filterFolders(nodes: Node[]): Node[] {
    return nodes
      .filter((node) => node.type === "folder")
      .map((node) => ({
        ...node,
        children: filterFolders(node.children || []),
      }));
  }

  const handleConfirm = () => {
    if (newName.trim() === "") return;
    const targetPath = selectedPath || "";
    onCreate(targetPath, newName.trim());
    onClose();
  };
</script>

<!-- Overlay (click outside or press ESC to close) -->
<div
  class="overlay"
  role="dialog"
  aria-modal="true"
  aria-label="Create new folder"
  tabindex="0"
  on:click={(e) => e.target === e.currentTarget && onClose()}
  on:keydown={(e) => {
    if (e.key === "Escape") {
      e.preventDefault();
      onClose();
    }
  }}
>
  <div class="modal" on:click|stopPropagation>
    <div class="modal-header">
      <h2>Create New Folder</h2>
      <IconButton icon={X} removeMargin={true} onClick={onClose} />
    </div>

    <div class="modal-body">
      <label>Select Directory:</label>
      <div class="folder-tree">
        <FolderTree
          nodes={filterFolders(tree)}
          selected={selectedPath}
          on:select={(e) => (selectedPath = e.detail.path)}
          isCollapsed={false}
        />
      </div>

      <label>New Folder Name:</label>
      <input
        type="text"
        placeholder="Enter folder name"
        bind:value={newName}
        on:keydown={(e) => e.key === "Enter" && handleConfirm()}
      />

      {#if selectedPath}
        <p class="selected-path">📁 Folder will be created inside: <b>{selectedPath}</b></p>
      {:else}
        <p class="selected-path muted">📁 Folder will be created at the root directory</p>
      {/if}
    </div>

    <div class="modal-footer">
      <button type="button" on:click={handleConfirm} disabled={!newName}>Create</button>
      <button type="button" class="cancel" on:click={onClose}>Cancel</button>
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
    animation: fadeIn 0.2s ease forwards;

    &:focus {
      outline: none;
    }
  }

  .modal {
    background: #1c1c1c;
    border-radius: 8px;
    width: 420px;
    max-height: 80vh;
    overflow-y: auto;
    padding: 0.1rem 1rem 1rem 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
    color: white;
    animation: scaleIn 0.25s ease forwards;
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .modal-body label {
    font-weight: 600;
    margin-top: 0.5rem;
  }

  .folder-tree {
    max-height: 220px;
    overflow-y: auto;
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 4px;
    padding: 0.3rem;
    margin-bottom: 0.8rem;
  }

  input {
    width: 60%;
    padding: 0.4rem 0.6rem;
    border-radius: 4px;
    border: none;
    background: #2a2a2a;
    color: white;
    margin-bottom: 0.5rem;

    &:focus {
      outline: 2px solid #e64d4d;
    }
  }

  .selected-path {
    font-size: 0.85rem;
    color: #ccc;
    margin-top: -0.3rem;
    &.muted {
      opacity: 0.7;
    }
  }

  .modal-footer {
    display: flex;
    justify-content: flex-end;
    gap: 0.5rem;

    button {
      padding: 0.5rem 1rem;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      font-weight: 500;
    }

    .cancel {
      background: #333;
      color: #ccc;
    }

    button:not(.cancel) {
      background: #e64d4d;
      color: white;
    }

    button:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }
  }

  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }

  @keyframes scaleIn {
    from { transform: scale(0.95); opacity: 0; }
    to { transform: scale(1); opacity: 1; }
  }
</style>
