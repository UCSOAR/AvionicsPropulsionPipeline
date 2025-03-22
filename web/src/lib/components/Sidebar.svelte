<script lang="ts">
  import type { Component } from "svelte";
  import IconButton from "./IconButton.svelte";
  import { ArrowUpToLine, MenuIcon, PanelLeftClose, PanelRightClose} from 'lucide-svelte';
  
  export let icon: Component | null = null;
  export let label: string;
  
  // State to track if sidebar is expanded or collapsed
  let isExpanded = false;
  
  // Toggle the sidebar state
  const toggleSidebar = () => {
    isExpanded = !isExpanded;
  };
  
  // Get the height of the TopBar to position the sidebar correctly
  import { onMount } from 'svelte';
  let topBarHeight = 0;
  
  onMount(() => {
    const topBar = document.querySelector('header.top-bar');
    if (topBar) {
      topBarHeight = topBar.getBoundingClientRect().height;
    } else {
      // Default fallback if TopBar not found
      topBarHeight = 60;
    }
  });
</script>

<!-- Toggle button that stays fixed -->
<button class="sidebar-toggle" style="top: {topBarHeight + 10}px" on:click={toggleSidebar}>
  {#if isExpanded}
    <PanelLeftClose size={18} />
  {:else}
    <PanelRightClose size={18} />
  {/if}
</button>

<!-- Overlay sidebar -->
<div class="sidebar-overlay" class:active={isExpanded} on:click={() => isExpanded = false}></div>

<aside class="side-bar" class:expanded={isExpanded} style="top: {topBarHeight}px">
  <div class="sidebar-content">
    <!-- Upload file button -->
    <button class="upload-button" on:click={() => {
      // Maintain the original upload functionality here
      // This is where you'd put your upload logic
    }}>
      <ArrowUpToLine size={20} />
      <span class="label">Upload File</span>
    </button>
    
    <!-- Files section header -->
    <div class="files-header">
      <h3>Files</h3>
      <button class="layout-button" on:click={toggleSidebar}>
        <PanelLeftClose size={20} />
      </button>
    </div>
    
    <!-- Empty file list section -->
    <div class="file-list">
      <!-- Files will be populated dynamically -->
    </div>
  </div>
</aside>

<style lang="scss">
  @use "../styles/variables.scss" as *;
  
  // Sidebar toggle button that stays fixed
  .sidebar-toggle {
    position: fixed;
    left: 10px;
    z-index: 999;
    background-color: #121212;
    color: white;
    border: none;
    border-radius: 50%;
    width: 40px;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
    
    &:hover {
      background-color: #222;
    }
  }
  
  // Dark overlay behind the sidebar
  .sidebar-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 998;
    opacity: 0;
    pointer-events: none;
    transition: opacity 0.3s ease;
    
    &.active {
      opacity: 1;
      pointer-events: auto;
    }
  }
  
  // Sidebar with fixed positioning
  .side-bar {
    position: fixed;
    left: -280px;
    width: 280px;
    background-color: #121212;
    z-index: 1000;
    box-sizing: border-box;
    transition: left 0.3s ease;
    box-shadow: 2px 0 10px rgba(0, 0, 0, 0.3);
    display: flex;
    flex-direction: column;
    bottom: 0; /* Extend to the bottom of the viewport */
    
    &.expanded {
      left: 0;
    }
  }
  
  .sidebar-content {
    flex: 1;
    overflow-y: auto; /* Allow scrolling of sidebar content if needed */
    padding: 1rem;
    display: flex;
    flex-direction: column;
  }
  
  // Upload button styling - preserving original transitions and hover effects
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
    margin-bottom: 1.5rem;
    cursor: pointer;
    
    $trans: all 0.12s ease;
    transition: $trans;
    
    :global(.lucide-icon),
    * {
      transition: $trans;
    }
    
    .label {
      font-size: 1rem;
      white-space: nowrap;
    }
    
    &:hover {
      background-color: $bg-color-highlighted;
      border-color: $txt-color-highlighted;
      
      :global(.lucide-icon) {
        stroke: $txt-color-highlighted;
      }
      
      span.label {
        color: $txt-color-highlighted;
      }
    }
  }
  
  // Files header section
  .files-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
    
    h3 {
      font-size: 1.25rem;
      font-weight: 600;
      color: white;
      margin: 0;
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
      transition: all 0.12s ease;
      border-radius: 4px;
      
      &:hover {
        background-color: rgba(255, 255, 255, 0.1);
      }
    }
  }
  
  // File list styling
  .file-list {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }
  
  button {
    border: 1px solid $outline-color-1;
    background-color: transparent;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.3rem;
    border-radius: $border-radius-1;
    padding: 0.6rem;
    
    $trans: all 0.12s ease;
    transition: $trans;
    
    :global(.lucide-icon),
    * {
      transition: $trans;
    }
    
    span.label {
      font-size: 0.95rem;
      white-space: nowrap;
    }
    
    &:hover {
      background-color: $bg-color-highlighted;
      border-color: $txt-color-highlighted;
      
      :global(.lucide-icon) {
        stroke: $txt-color-highlighted;
      }
      
      span.label {
        color: $txt-color-highlighted;
      }
    }
  }
</style>