<script lang="ts">
  import { ChevronDown } from "@lucide/svelte";

  export let id: string;
  export let label: string | null = null;
  export let options: string[];

  if (!options) {
    throw new Error("At least one option must be provided.");
  }
</script>

<div class="container" {id}>
  {#if label}
    <label for={id}>{label}</label>
  {/if}
  <button>
    <span>{options[0]}</span>
    <ChevronDown />
  </button>
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.container {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;

    label {
      white-space: nowrap;
    }

    button {
      background-color: $bg-color-4;
      border-radius: $border-radius-1;
      padding: 0.55rem;
      font-size: 0.95rem;
      display: flex;
      align-items: center;
      gap: 0.8rem;

      :global(.lucide-icon) {
        $size: 1rem;

        width: $size;
        height: $size;
        stroke: $txt-color-1;
        opacity: 0.5;
        transition: opacity 0.15s ease;
        margin-left: auto;
      }

      &:hover {
        :global(.lucide-icon) {
          opacity: 1;
        }
      }
    }
  }
</style>
