<script lang="ts">
  import type { Component } from "svelte";

  export let icon: Component | null = null;
  export let label: string | null = null;
  export let isDisabled: boolean = false;
  export let onClick: () => void;
</script>

<button onclick={onClick} disabled={isDisabled}>
  {#if icon}
    <svelte:component this={icon} />
  {/if}
  {#if label}
    <span class="label">{label}</span>
  {/if}
</button>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  button {
    border: 1px solid $outline-color-1;
    background-color: transparent;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.4rem;
    border-radius: $border-radius-1;
    padding: 0.6rem;

    $trans: all 0.12s ease;

    transition: $trans;

    :global(.lucide-icon) {
      $size: 1.1rem;

      width: $size;
      height: $size;
    }

    :global(.lucide-icon),
    * {
      transition: $trans;
    }

    span.label {
      font-size: 0.95rem;
    }

    &:not([disabled]):hover {
      background-color: $bg-color-highlighted;
      border-color: $txt-color-highlighted;

      :global(.lucide-icon) {
        stroke: $txt-color-highlighted;
      }

      span.label {
        color: $txt-color-highlighted;
      }
    }

    &:disabled {
      background-color: $bg-color-2;
      border-color: $outline-color-1;
      cursor: not-allowed;

      :global(.lucide-icon) {
        stroke: $outline-color-1;
      }

      span.label {
        color: $outline-color-1;
      }
    }
  }
</style>
