<script lang="ts"></script>

<header class="top-bar">
  <div class="logo">
    <img draggable={false} src="/soar-logo.svg" alt="SOAR Logo" />
  </div>
  <nav class="nav-links">
    <a href="/">TAC</a>
    <a href="/">HAC</a>
    <a href="/">Mapleleaf</a>
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
    margin-right: auto; /* Moves logo to the left */
  }

  .nav-links {
    display: flex;
    justify-content: center; /* Centers links inside */
    flex-grow: 1; /* Allows nav-links to take available space */
    gap: 2rem;
  }

  .nav-links a {
    text-decoration: none;
    color: $txt-color-2; /* Light text */
    font-size: 1.1rem;
    transition: color 0.2s ease;
    outline: none;
  }

  .logo img {
    height: 3rem; /* Restores original size */
    width: auto; /* Keeps aspect ratio */
  }

  .nav-links a:hover {
    color: $txt-color-highlighted; /* Change color on hover */
  }
</style>
