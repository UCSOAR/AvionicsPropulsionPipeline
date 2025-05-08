<script lang="ts">
  import IconButton from "./IconButton.svelte";
  import { onMount } from "svelte";
  import { Expand, Shrink } from "@lucide/svelte";
  import { exitFullscreen, requestFullscreen } from "$lib/utils/fullscreen";

  export let targetElement: HTMLElement;
  export let onChange: (isFullscreen: boolean) => void;
  let isFullscreen = false;

  function toggleFullscreen() {
    if (isFullscreen) {
      exitFullscreen();
    } else {
      requestFullscreen(targetElement);
    }
  }

  function onFullscreenChange() {
    isFullscreen = document.fullscreenElement === targetElement;
    onChange(isFullscreen);
  }

  onMount(() => {
    document.addEventListener("fullscreenchange", onFullscreenChange);

    return () => {
      document.removeEventListener("fullscreenchange", onFullscreenChange);
    };
  });
</script>

{#if isFullscreen}
  <IconButton icon={Shrink} onClick={toggleFullscreen} />
{:else}
  <IconButton icon={Expand} onClick={toggleFullscreen} />
{/if}
