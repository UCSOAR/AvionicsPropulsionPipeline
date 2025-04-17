<script lang="ts">
  import "$lib/styles/global.scss";
  import type { SelectedFile } from "$lib/models/selectedFile";
  import { goto } from '$app/navigation';

  function goToTAC() {
    goto('/tac');
  }

  let selectedFile: SelectedFile | undefined = undefined;
</script>

<header class="front-page">
  <div class="logo">
    <img draggable={false} src="/soar-logo.svg" alt="SOAR Logo" />
  </div>
</header>

<div>
  <div class="header-text">
    <h1>Propulsion Pipeline</h1>
    <p>
      Automating the post-processing of propulsion test data and<br />
      archiving results online for analysis
    </p>
  </div>

  <div class="card-container">
    <div class="card" on:click={goToTAC}>
      <h2>TAC</h2>
      <p>definition:</p>
    </div>

    <div class="card">
      <h2>HAC</h2>
      <p>definition:</p>
    </div>

    <div class="card">
      <h2>Mapleleaf</h2>
      <p>definition:</p>
    </div>
  </div>
</div>

<style lang="scss">
  @use "../lib/styles/variables.scss" as *;

  body {
    background-color: #000;
    margin: 0;
    padding: 0;
  }

  .logo {
    position: absolute;
    top: 2rem;
    left: 2rem;
  }

  .logo img {
    height: 4rem;
    width: auto;
  }

  .header-text {
    color: white;
    text-align: center;
    margin-top: 3rem;

    h1 {
      color: red;
      font-size: 2.5rem;
      margin: 0.5rem 0;
    }

    p {
      color: white;
      font-size: 1rem;
      margin: 0.5rem;
    }
  }

  .card-container {
    display: flex;
    justify-content: center;
    gap: 4rem;
    margin-top: 8rem;
    padding: 0 2rem;
    flex-wrap: wrap;
  }

  .card {
    background: linear-gradient(to bottom, #3b1c1c, #1a0a0a);
    border-radius: 1rem;
    padding: 2rem;
    width: 12rem;
    height: 12rem;
    color: white;
    text-align: center;
    box-shadow: 0 0 20px rgba(255, 0, 0, 0.2);
    transition: transform 0.2s ease;
    cursor: pointer;

    &:hover {
      transform: scale(1.05);
    }

    h2 {
      font-size: 1.5rem;
      margin-bottom: 1rem;
    }

    p {
      font-size: 0.9rem;
      color: #ddd;
    }
  }
</style>
