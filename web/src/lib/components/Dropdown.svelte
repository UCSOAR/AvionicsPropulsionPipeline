<script lang="ts">
  import { Check, CheckCheck, ChevronDown, ChevronUp } from "@lucide/svelte";
  import { onMount } from "svelte";

  export let id: string;
  export let label: string | null = null;
  export let options: string[];
  export let onChange: (optionIndex: number) => void;

  if (!options) {
    throw new Error("At least one option must be provided.");
  }

  let selectedOptionIndex = 0;
  let isOptionsVisible = false;

  const handleSelect = (index: number) => {
    selectedOptionIndex = index;
    isOptionsVisible = false;
    onChange(index);
  };

  const toggleOptions = () => {
    isOptionsVisible = !isOptionsVisible;
  };

  onMount(() => {
    const handleClickOutside = (event: MouseEvent) => {
      const target = event.target as HTMLElement;
      const container = document.querySelector(`div.container[id="${id}"]`);

      if (!container?.contains(target)) {
        isOptionsVisible = false;
      }
    };

    document.addEventListener("click", handleClickOutside);
    return () => document.removeEventListener("click", handleClickOutside);
  });
</script>

<div class="container" {id}>
  {#if label}
    <label for={id}>{label}</label>
  {/if}
  <button class="dropdown-button" on:click={toggleOptions}>
    <span>{options[selectedOptionIndex]}</span>
    {#if !isOptionsVisible}
      <ChevronDown />
    {:else}
      <ChevronUp />
    {/if}
  </button>
  <ul class="options-container" class:visible={isOptionsVisible}>
    {#each options as option, index}
      <li>
        <button
          type="button"
          class:selected={index === selectedOptionIndex}
          on:click={() => handleSelect(index)}
        >
          {option}
          {#if index === selectedOptionIndex}
            <Check />
          {/if}
        </button>
      </li>
    {/each}
  </ul>
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.container {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;
    position: relative;

    label {
      white-space: nowrap;
    }

    button.dropdown-button {
      background-color: $bg-color-6;
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

    ul.options-container {
      display: none;
      list-style: none;
      position: absolute;
      box-sizing: border-box;
      background-color: $bg-color-6;
      z-index: 100;
      padding: 0.5rem;
      gap: 0.3rem;
      margin: 0;
      margin-top: 0.5rem;
      border-radius: $border-radius-1;

      & > li {
        width: 100%;

        * {
          transition: all 0.1s ease;
        }

        & > button {
          display: flex;
          flex-direction: row;
          border-radius: $border-radius-1 * 0.7;
          background: none;
          padding: 0.6rem;
          width: 100%;

          :global(.lucide-icon) {
            width: 1rem;
            height: auto;
            stroke: $txt-color-1;
            margin-left: auto;
          }

          &:hover {
            background-color: $bg-color-highlighted;
            color: $txt-color-highlighted;

            :global(.lucide-icon) {
              stroke: $txt-color-highlighted;
            }
          }
        }
      }

      &.visible {
        display: flex;
        flex-direction: column;
        top: 100%;
        width: 100%;
        left: 0;
      }
    }
  }
</style>
