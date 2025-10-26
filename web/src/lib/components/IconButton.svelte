<script lang="ts">
  import type { Component } from "svelte";

  export let icon: Component | null = null;
  export let label: string | null = null;
  export let isDisabled: boolean = false;
  export let onClick: () => void;
  export let toggle: boolean = false;
  export let removeMargin: boolean = false;

  let internalToggle = false;

  function handleClick() {

    if (toggle) {
      internalToggle = !internalToggle;
    }

    if (onClick) onClick();
  }

  

</script>

<button class ="remove-margin" onclick={handleClick} disabled={isDisabled} class:active={internalToggle} class:remove-margin={removeMargin}>
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

    &.active {
      background-color: $bg-color-highlighted;
      border-color: $txt-color-highlighted;

      :global(.lucide-icon) {
        stroke: $txt-color-highlighted;
      }
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

  &.remove-margin {
    border: 0;
    padding-right: 0;

    &:hover,
    &.active {
      background-color: transparent;
      border-color: transparent;

      :global(.lucide-icon) {
        stroke: $txt-color-1; 
      }

      span.label {
        color: $txt-color-1;
      }
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
