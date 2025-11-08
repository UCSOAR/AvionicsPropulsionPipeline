<script context="module" lang="ts">
  // ✅ Type definition for a tree node
  export type Node = {
    name: string;             // Display name (e.g. "Documents")
    path: string;             // Full relative path (e.g. "Documents/notes")
    type: "folder" | "file";  // Node type
    children?: Node[];        // Subfolders or files
    metadata?: any;           // Optional file metadata
  };
</script>

<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { ChevronRight, ChevronDown, Folder as FolderClosed, FolderOpen, File } from "lucide-svelte";

  // Props
  export let nodes: Node[] = [];
  export let selected: string = "";
  export let isCollapsed: boolean = false;
  export let open: Set<string> = new Set(); 

  // Dispatcher for events
  const dispatch = createEventDispatcher();

  // --- Helpers ---
  const toggle = (path: string) => {
    const newOpen = new Set(open);
    if (newOpen.has(path)) newOpen.delete(path);
    else newOpen.add(path);
    open = newOpen;
  };

  const isOpen = (path: string) => open.has(path);

  // Dispatch selection event
  const handleSelect = (node: Node) => {
    console.log("📁 Dispatching select:", node);
    dispatch("select", {
      path: node.path,
      metadata: node.metadata,
      type: node.type,
    });
  };
</script>

<ul class="tree">
  {#each nodes as node (node.path)}
    <li>
      {#if node.type === "folder"}
        <!-- Folder Button -->
        <button
          type="button"
          class="folder-row {selected === node.path ? 'selected' : ''}"
          on:click={(e) => {
            e.stopPropagation();
            toggle(node.path);
            handleSelect(node);
          }}
          title={node.name}
        >
          {#if isOpen(node.path)}
            <ChevronDown size={14} />
            <FolderOpen size={14} />
          {:else}
            <ChevronRight size={14} />
            <FolderClosed size={14} />
          {/if}

          {#if !isCollapsed}
            <span class="label">{node.name}</span>
          {/if}
        </button>

        <!-- Recursive child rendering -->
        {#if isOpen(node.path) && node.children?.length}
          <div class="children">
            <svelte:self
              nodes={node.children}
              selected={selected}
              isCollapsed={isCollapsed}
              on:select={(e) => dispatch("select", e.detail)}
              open={open}
            />
          </div>
        {/if}

      {:else}
        <!-- File Button -->
        <button
          type="button"
          class="file-row {selected === node.path ? 'selected' : ''}"
          on:click={(e) => {
            e.stopPropagation();
            handleSelect(node);
          }}
          title={node.name}
        >
          <File size={14} />
          {#if !isCollapsed}
            <span class="label">{node.name}</span>
          {/if}
        </button>
      {/if}
    </li>
  {/each}
</ul>

<style lang="scss">
  ul.tree {
    list-style: none;
    padding-left: 0;
    margin: 0;
  }

  li {
    margin: 2px 0;
  }

  button.folder-row,
  button.file-row {
    display: flex;
    align-items: center;
    gap: 0.45rem;
    padding: 0.35rem 0.55rem;
    color: white;
    border-radius: 4px;
    cursor: pointer;
    user-select: none;
    background: transparent;
    border: none;
    width: 100%;
    text-align: left;
    transition: background 0.15s ease, color 0.15s ease;

    &:hover {
      background: rgba(255, 255, 255, 0.08);
    }

    &.selected {
      background: rgba(255, 255, 255, 0.15);
      color: #e64d4d;

      :global(svg) {
        stroke: #e64d4d;
      }
    }
  }

  .label {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .children {
    margin-left: 1rem;
    padding-left: 0.4rem;
    border-left: 1px solid rgba(255, 255, 255, 0.05);
    overflow: hidden;
    transition: all 0.2s ease-in-out;
  }

  .folder-row {
    font-weight: 500;

    :global(svg) {
      opacity: 0.9;
    }
  }

  .file-row {
    font-weight: 400;
  }
</style>
