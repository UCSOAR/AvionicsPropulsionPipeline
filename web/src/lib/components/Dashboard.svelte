<script lang="ts">
  import Dropdown from "./Dropdown.svelte";
  import Input from "./Input.svelte";
  import IconButton from "./IconButton.svelte";
  import { browser } from "$app/environment";
  import { writable } from "svelte/store";
  import { onMount } from "svelte";
  import { fetchStaticFireColumns } from "$lib/utils/getStaticFireColumns";
  import { numericRegex } from "$lib/utils/regexps";
  import type { Config, Data, Layout } from "plotly.js";
  import type { PostStaticFireColumnsRequest } from "$lib/models/dashboardModels";
  import type { SelectedFile } from "$lib/models/selectedFile";
  import {
    Loader2,
    MessageCircleWarningIcon,
    RefreshCcw,
  } from "@lucide/svelte";

  export let selectedFile: SelectedFile;

  let plotlyChartDiv: HTMLDivElement;
  let selectedXColumnIndex = writable(0);
  let selectedYColumnIndex = writable(0);
  let startRow = 0;
  let numRows = 0;
  let totalRows = 0;
  let isLoadingPlotly = false;
  let plotError = "";

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

  const safeParseInt = (value: string) => {
    const parsedValue = parseInt(value, 10);
    return isNaN(parsedValue) ? 0 : parsedValue;
  };

  const loadPlotly = async (
    fetchData?: () => Promise<Partial<Data>[] | null>
  ) => {
    if (isLoadingPlotly) {
      return;
    }

    isLoadingPlotly = true;

    const data = fetchData !== undefined ? await fetchData() : [];

    if (!data) {
      plotError = "Failed to fetch data.";
    }
    const Plotly = await import("plotly.js-dist-min");
    await Plotly.newPlot(plotlyChartDiv, data ?? [], layout, config);

    isLoadingPlotly = false;
  };

  const refreshPlotly = async () => {
    if (!selectedFile) {
      return;
    }

    plotError = "";

    const xColumnName =
      selectedFile.metadata.xColumnNames[$selectedXColumnIndex];
    const yColumnName =
      selectedFile.metadata.yColumnNames[$selectedYColumnIndex];

    // Test request for now
    const req: PostStaticFireColumnsRequest = {
      name: selectedFile.name,
      startRow,
      numRows,
      xColumnNames: [xColumnName],
      yColumnNames: [yColumnName],
    };

    await loadPlotly(async () => {
      const res = await fetchStaticFireColumns(req);

      if (!res) {
        return null;
      }

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

  selectedXColumnIndex.subscribe(refreshPlotly);
  selectedYColumnIndex.subscribe(refreshPlotly);

  onMount(loadPlotly);
</script>

<div class="container">
  <div class="content-header">
    <div class="title">
      <h1>Dashboard for <i>{selectedFile.name}</i></h1>
      <p>
        Visualizing data for <i
          >{selectedFile.metadata.xColumnNames[$selectedXColumnIndex]}</i
        >
        and
        <i>{selectedFile.metadata.yColumnNames[$selectedYColumnIndex]}</i>
      </p>
    </div>
    <div class="data-select">
      <div class="column-select">
        <Dropdown
          onChange={(index) => ($selectedXColumnIndex = index)}
          isDisabled={isLoadingPlotly}
          label="X Column"
          id="x-column"
          options={selectedFile.metadata.xColumnNames}
        />
        <Dropdown
          onChange={(index) => ($selectedYColumnIndex = index)}
          isDisabled={isLoadingPlotly}
          label="Y Column"
          id="y-column"
          options={selectedFile.metadata.yColumnNames}
        />
      </div>
      <div class="row-select">
        <Input
          id="start-row"
          placeholder="0"
          isDisabled={isLoadingPlotly}
          label="Start Row"
          regex={numericRegex}
          onChange={(value) => (startRow = safeParseInt(value))}
        />
        <Input
          id="num-rows"
          value= {null}
          placeholder= {selectedFile.metadata.totalRows.toString()}
          isDisabled={isLoadingPlotly}
          label= {`Row Count`}
          regex={numericRegex}
          onChange={(value) => (numRows = safeParseInt(value))}
        />
      </div>
    </div>
  </div>
  <div class="content-container">
    <div class="chart-pod pod">
      <div class="title-container">
        <h2>Static Fire Chart</h2>
        <IconButton icon={RefreshCcw} onClick={refreshPlotly} />
      </div>
      <div class="chart-wrapper">
        <div
          class="loading-overlay"
          class:hidden={!isLoadingPlotly && !plotError}
        >
          {#if plotError}
            <div>
              <MessageCircleWarningIcon />
              <b>{plotError}</b>
            </div>
          {:else}
            <Loader2 />
          {/if}
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
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  .container {
    flex-grow: 1; // This makes the dashboard expand to fill the remaining space
    padding: 1rem;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
  }

  div.content-header {
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    gap: 2rem;

    div.title {
      margin-right: auto;
    }

    div.data-select {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      width: 100%;
      gap: 1rem;
      margin-bottom: 0.8rem;

      & > div {
        gap: 1.5rem;
        flex-grow: 1;
        display: flex;
        flex-direction: row;
      }

      & > div.row-select {
        justify-content: flex-end;
      }
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

      div.title-container {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 0.7rem;

        h2 {
          margin: 0;
        }
      }

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

          & > div {
            display: flex;
            flex-direction: column;
            gap: 0.6rem;
            justify-content: center;
            align-items: center;

            b {
              color: $txt-color-highlighted;
              font-size: 1.2rem;
            }

            :global(.lucide-icon) {
              width: 3rem;
              height: auto;
              stroke: $txt-color-highlighted;
            }
          }
        }
      }
    }
  }
</style>
