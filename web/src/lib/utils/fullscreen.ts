// utils/fullscreen.ts
export function requestFullscreen(el: HTMLElement) {
    if (el.requestFullscreen) return el.requestFullscreen();
    if ((el as any).webkitRequestFullscreen) return (el as any).webkitRequestFullscreen();
    if ((el as any).mozRequestFullScreen) return (el as any).mozRequestFullScreen();
    if ((el as any).msRequestFullscreen) return (el as any).msRequestFullscreen();
  }
 
  export function exitFullscreen() {
    if (document.exitFullscreen) return document.exitFullscreen();
    if ((document as any).webkitExitFullscreen) return (document as any).webkitExitFullscreen();
    if ((document as any).mozCancelFullScreen) return (document as any).mozCancelFullScreen();
    if ((document as any).msExitFullscreen) return (document as any).msExitFullscreen();
  }