<script lang="ts">
  import type { Component } from 'svelte';

  export let icon: Component | null = null;
  export let label: string = '';
</script>

<header class="side-bar">
  <div class="logo"></div>

  <button aria-label={label}>
    {#if icon}
      <svelte:component this={icon} />
    {/if}
    <span class="label">{label}</span>
  </button>

  <div class="divider"></div>

  <div class="file-section">
    <h3 class="title">Files      
    </h3>
  </div>
</header>

<style scoped lang="scss">
  @use '../styles/variables.scss' as *;

  .side-bar {
    width: 250px;
    background-color: #121212;
    height: 100vh;
    padding: 1rem;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    position: fixed;
    top: 0;
    left: 0;
  }


  button {
    display: flex;
    align-items: center;
    gap: 0.3rem;
    padding: 0.6rem;
    background-color: transparent;
    border: 1px solid $outline-color-1;
    border-radius: $border-radius-1;
    transition: all 0.12s ease;

    span.label {
      font-size: 0.95rem;
      transition: color 0.12s ease;
    }

    &:hover {
      background-color: $bg-color-highlighted;
      border-color: $txt-color-highlighted;

      :global(.lucide-icon) {
        stroke: $txt-color-highlighted;
      }

      .label {
        color: $txt-color-highlighted;
      }
    }
  }

  .divider {
    height: 1px;
    background-color: #333;
    margin: 1rem 0;
  }

  .file-section {
    padding-left: 0.25rem;
  }

  .title {
    color: white;
    font-size: 1rem;
    font-weight: 600;
    margin-bottom: 0.5rem;
  }
</style>
