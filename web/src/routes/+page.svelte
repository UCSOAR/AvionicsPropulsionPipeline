<script lang="ts">
  import "$lib/styles/global.scss";
  import "$lib/components/IconButton.svelte";
  import type { SelectedFile } from "$lib/models/selectedFile";
  import IconButton from "$lib/components/IconButton.svelte";
  import { LucideChartBar } from "@lucide/svelte";
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

  <div class="button-container">
    <IconButton icon={LucideChartBar} label="TAC" onClick={goToTAC} />
    <IconButton label="HAC" onClick={() => {}} />
    <IconButton label="Mapleleaf" onClick={() => {}} />
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


  .button-container {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 20rem;
    margin-top: 15rem;
  }
</style>