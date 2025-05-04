<script lang="ts">
  import { goto } from "$app/navigation";

  function handleLogoClick() {
    goto("/");
  }
</script>

<header class="top-bar">
  <div
    class="logo"
    role="button"
    tabindex="0"
    on:click={handleLogoClick}
    on:keydown={handleLogoClick}
  >
    <img draggable={false} src="/soar-logo.svg" alt="SOAR Logo" />
  </div>

  <nav class="nav-links">
    <button>TAC</button>
    <button>HAC</button>
    <button>Mapleleaf</button>
  </nav>
</header>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  .top-bar {
    display: flex;
    align-items: center;
    gap: 1.3rem;
    padding: 0.7rem;
    height: auto;
    background-color: $bg-color-2;
    border-bottom: 1px solid $outline-color-1;
  }

  .logo {
    background: none;
    border: none;
    padding: 0;
    margin-right: auto;
    cursor: pointer;
    border-radius: 0.5rem;

    img {
      height: 3rem;
      width: auto;
      display: block;
    }
  }

  .nav-links {
    display: flex;
    justify-content: center;
    flex-grow: 1;
    gap: 2rem;
  }

  .nav-links button {
    background: none;
    border: none;
    color: $txt-color-2;
    font-size: 1.1rem;
    transition: color 0.2s ease;
    cursor: pointer;
    padding: 0;
  }

  .nav-links button:hover {
    color: $txt-color-highlighted;
  }

  .nav-links button:focus {
    outline: none;
  }
</style>
