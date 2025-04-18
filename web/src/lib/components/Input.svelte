<script lang="ts">
  import type { HTMLInputTypeAttribute } from "svelte/elements";

  export let id: string;
  export let value: number | null = 0;
  export let placeholder: string;
  export let label: string | null = null;
  export let type: HTMLInputTypeAttribute = "text";
  export let isDisabled: boolean = false;
  export let regex: RegExp | null = null;
  export let onChange: (value: string) => void;

  let isError: boolean = false;

  const handleInput = (event: Event) => {
    const target = event.target as HTMLInputElement;
    const value = target.value;

    if (value.length === 0) {
      isError = false;
    } else if (regex) {
      isError = !regex.test(value);
    }

    onChange(value);
  };
</script>

<div class="container" {id}>
  {#if label}
    <label for={id}>{label}</label>
  {/if}
  <input
    disabled={isDisabled}
    {type}
    {id}
    {value}
    class="input-field"
    class:error={isError}
    on:input={handleInput}
    {placeholder}
  />
</div>

<style lang="scss">
  @use "../styles/variables.scss" as *;

  div.container {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;

    input {
      background-color: $bg-color-6;
      border-radius: $border-radius-1;
      border: 1px solid transparent;
      padding: 0.55rem;
      font-size: 0.95rem;
      display: flex;
      align-items: center;
      gap: 0.8rem;
      transition: all 0.06s ease;
      min-width: 4rem;
      max-width: 7rem;
      transition: all 0.12s ease;

      &:disabled {
        cursor: not-allowed;
      }

      &.error {
        background-color: $bg-color-highlighted;
        color: $txt-color-highlighted;
        border-color: $txt-color-highlighted;
      }
    }
  }
</style>
