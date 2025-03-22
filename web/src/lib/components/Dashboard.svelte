<script lang="ts">
  import Dropdown from "./Dropdown.svelte";
  import { browser } from "$app/environment";
  import { onMount } from "svelte";
  import { Loader2 } from "@lucide/svelte";
  import { fetchStaticFireColumns } from "$lib/utils/getStaticFireColumns";
  import type { PreviewMetadata } from "$lib/models/cacheTreeModels";
  import type { Config, Data, Layout } from "plotly.js";
  import type { PostStaticFireColumnsRequest } from "$lib/models/dashboardModels";

  type SelectedFile = {
    name: string;
    metadata: PreviewMetadata;
  };

  export let selectedFile: SelectedFile | null = null;

  const style = {
    margin: 50,
    bgColor: "#1f1f1f",
    txtColor: "#e1e1e1",
    themeColor: "#dc2626",
  };
  const config: Partial<Config> = { responsive: true };
  const layout: Partial<Layout> = {
    autosize: true,
    margin: {
      l: style.margin,
      r: style.margin,
      t: style.margin,
      b: style.margin,
    },
    paper_bgcolor: style.bgColor,
    plot_bgcolor: style.bgColor,
    font: {
      family: "Inter",
      color: "white",
    },
    xaxis: { color: style.txtColor },
    yaxis: { color: style.txtColor },
  };

  let selectedXColumnIndex = 0;
  let selectedYColumnIndex = 0;
  let plotlyChartDiv: HTMLDivElement;
  let isLoadingPlotly = false;

  const loadPlotly = async (fetchData?: () => Promise<Partial<Data>[]>) => {
    if (isLoadingPlotly) {
      return;
    }

    isLoadingPlotly = true;

    const data = fetchData !== undefined ? await fetchData() : [];
    const Plotly = await import("plotly.js-dist-min");
    await Plotly.newPlot(plotlyChartDiv, data, layout, config);

    isLoadingPlotly = false;
  };

  onMount(loadPlotly);

  // Update chart data when selected columns change
  $: {
    if (selectedFile && browser) {
      const updatePlotly = async () => {
        const xColumnName =
          selectedFile.metadata.xColumnNames[selectedXColumnIndex];
        const yColumnName =
          selectedFile.metadata.yColumnNames[selectedYColumnIndex];

        // Test request for now
        const req: PostStaticFireColumnsRequest = {
          name: selectedFile.name,
          startRow: 0,
          numRows: 40000,
          xColumnNames: [xColumnName],
          yColumnNames: [yColumnName],
        };

        await loadPlotly(async () => {
          const res = await fetchStaticFireColumns(req);

          if (!res) {
            return [];
          }
          console.log(res.yColumns[yColumnName].rows.length);

          const data: Partial<Data> = {
            x: res.xColumns[xColumnName].rows,
            y: res.yColumns[yColumnName].rows,
            type: "scattergl",
            mode: "lines",
            line: { color: style.themeColor },
          };

          return [data];
        });
      };

      updatePlotly();
    }
  }
</script>

<div class="container">
  {#if selectedFile}
    <div class="content-header">
      <div class="title">
        <h1>Dashboard for <i>{selectedFile.name}</i></h1>
        <p>
          Visualizing data for <i
            >{selectedFile.metadata.xColumnNames[selectedXColumnIndex]}</i
          >
          and
          <i>{selectedFile.metadata.yColumnNames[selectedYColumnIndex]}</i>
        </p>
      </div>
      <div class="column-select">
        <Dropdown
          onChange={(index) => (selectedXColumnIndex = index)}
          isDisabled={isLoadingPlotly}
          label="X Column"
          id="x-column"
          options={selectedFile.metadata.xColumnNames}
        />
        <Dropdown
          onChange={(index) => (selectedYColumnIndex = index)}
          isDisabled={isLoadingPlotly}
          label="Y Column"
          id="y-column"
          options={selectedFile.metadata.yColumnNames}
        />
      </div>
    </div>
    <div class="content-container">
      <div class="chart-pod pod">
        <h2>Static Fire Chart</h2>
        <div class="chart-wrapper">
          <div class="loading-overlay" class:hidden={!isLoadingPlotly}>
            <Loader2 class="spinner" />
          </div>
          <div bind:this={plotlyChartDiv} class="chart"></div>
        </div>
      </div>
      <div class="value-pods">
        <div class="min-val-pod pod">
          <label for="min-val">Minimum Value</label>
          <div class="value" id="min-val">0.00</div>
        </div>
        <div class="max-val-pod pod">
          <label for="max-val">Maximum Value</label>
          <div class="value" id="max-val">0.00</div>
        </div>
        <div class="avg-val-pod pod">
          <label for="avg-val">Average Value</label>
          <div class="value" id="avg-val">0.00</div>
        </div>
      </div>
    </div>
  {:else}
    <h1>Dashboard</h1>
    <div class="message-container">
      <p>Select a file from the sidebar to view its data.</p>
    </div>
  {/if}
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.content-header {
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    gap: 2rem;

    div.title {
      margin-right: auto;
    }

    div.column-select {
      display: flex;
      flex-direction: row;
      gap: 1.5rem;
      width: 25%;
      min-width: 9rem;
      max-width: 13rem;
    }
  }

  div.pod {
    background-color: $bg-color-4;
    border: 1px solid $outline-color-1;
    border-radius: $border-radius-1;
    padding: 1rem;

    h2 {
      margin: 0;
      margin-bottom: 0.5em;
    }

    div.value {
      margin-top: 0.5rem;
      font-size: 1.5rem;
      font-weight: bold;
    }
  }

  div.content-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding-bottom: 1rem;

    div.value-pods {
      display: flex;
      flex-direction: row;
      gap: 1rem;

      div.pod {
        flex-grow: 1;
      }
    }

    div.chart-pod {
      flex-grow: 1;

      div.chart-wrapper {
        border-radius: $border-radius-1;
        overflow: hidden;
        position: relative;

        div.loading-overlay {
          position: absolute;
          top: 0;
          left: 0;
          width: 100%;
          height: 100%;
          display: flex;
          justify-content: center;
          align-items: center;
          background-color: rgba(0, 0, 0, 0.5);
          z-index: 1;

          :global(.lucide-icon) {
            width: 3.3rem;
            height: auto;
          }

          &.hidden {
            display: none;
          }
        }
      }
    }
  }

  div.container {
    padding: 1.3rem;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    gap: 1rem;

    h1 {
      margin: 0;
      margin-bottom: 0.3rem;

      i {
        color: $txt-color-1;
      }
    }

    div.message-container {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-grow: 1;
      text-align: center;
    }
  }
</style>
