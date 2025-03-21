<script lang="ts">
  import type { Component } from "svelte";

  export let icon: Component | null = null;
  export let label: string;
</script>

<header class="side-bar">
  <div class = "logo">
  </div>


<button>
  {#if icon}
    <svelte:component this={icon} />
  {/if}
  <span class="label">{label}</span>
</button>

<div class="divider"></div>

<div class ="file-section"></div>
  <h3 class="title">Files</h3>

</header>

<style scoped lang="scss">


  @use "../styles/variables.scss" as *;

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
 
  // make divider lines and files here
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
  .divider {
  height: 1px;
  background-color: #333;
  margin: 1rem 0;
  }

  .title {
    color: white;
    font-size: 1rem;
    font-weight: 600;
    margin-bottom: 0.5rem;
    padding-left: 0.25rem;
  }



</style>