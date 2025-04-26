<script lang="ts">
    import { onMount, onDestroy } from "svelte";
   
    export let targetElement: HTMLElement;
   
    let isFullscreen = false;
   
    function enterFullscreen() {
      if (targetElement?.requestFullscreen) targetElement.requestFullscreen();
      else if ((targetElement as any)?.webkitRequestFullscreen) (targetElement as any).webkitRequestFullscreen();
      else if ((targetElement as any)?.mozRequestFullScreen) (targetElement as any).mozRequestFullScreen();
      else if ((targetElement as any)?.msRequestFullscreen) (targetElement as any).msRequestFullscreen();
    }
   
    function exitFullscreen() {
      if (document.exitFullscreen) document.exitFullscreen();
      else if ((document as any)?.webkitExitFullscreen) (document as any).webkitExitFullscreen();
      else if ((document as any)?.mozCancelFullScreen) (document as any).mozCancelFullScreen();
      else if ((document as any)?.msExitFullscreen) (document as any).msExitFullscreen();
    }
   
    function toggleFullscreen() {
      isFullscreen ? exitFullscreen() : enterFullscreen();
    }
   
    function handleFullscreenChange() {
      isFullscreen = document.fullscreenElement === targetElement;
    }
   
    onMount(() => {
      document.addEventListener("fullscreenchange", handleFullscreenChange);
    });
   
    onDestroy(() => {
      document.removeEventListener("fullscreenchange", handleFullscreenChange);
    });
  </script>
   
  <style>
    button {
      background: none;
      border: 1px solid white;
      padding: 0.5rem 1rem;
      border-radius: 0.5rem;
      color: white;
      cursor: pointer;
      font-size: 0.9rem;
      transition: background 0.2s ease;
    }
   
    button:hover {
      background: rgba(255, 255, 255, 0.1);
    }
  </style>
   
  <button on:click={toggleFullscreen}>
    {isFullscreen ? "Exit Fullscreen" : "Go Fullscreen"}
  </button>