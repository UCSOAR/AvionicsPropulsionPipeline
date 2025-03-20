<script lang="ts">
  import type { Component } from "svelte";

 
</script>

<header class="top-bar">
  <div class="logo">
    <img src="/soar-logo.svg" alt="" />
  </div>

  <nav class="nav-links">
    <a href="#">pyTac</a>
    <a href="#">pyHac</a>
    <a href="#">mapleleaf</a>
  </nav>
</header>

  

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;
    
  .top-bar {
  display: flex;
  align-items: center;
  justify-content: center; /* Keep the navigation centered */
  padding: 0 1rem;
  height: 60px;
  background-color: black;
  position: relative; /* Allows absolute positioning for the logo */
}

.logo {
  position: absolute;
  left: 1rem; /* Move logo to the left */
}
  .nav-links {
    display: flex;
    gap: 2rem;
  }

  .nav-links a {
    text-decoration: none;
    color: white; /* Light text */
    font-size: 1.2rem;
    font-weight: 500;
    transition: color 0.2s ease;
  }

  .logo img {
  height: 50px; /* Restores original size */
  width: auto; /* Keeps aspect ratio */
}

  .nav-links a:hover {
    color: red; /* Change color on hover */
  }
  
</style> 
