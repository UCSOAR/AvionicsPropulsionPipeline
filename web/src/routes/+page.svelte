<script lang="ts">
  import "$lib/styles/global.scss";
  import "$lib/components/IconButton.svelte";
  import Dashboard from "$lib/components/Dashboard.svelte";
  import TopBar from "$lib/components/TopBar.svelte";
</script>

<main>
  <TopBar />
  <Dashboard
    selectedFile={{
      name: "RHT_2023-09-21-2-33_PM (1)(1)",
      metadata: {
        operator: "Test Operator",
        resultTimestamp: {
          date: "2021-09-01",
          time: "12:00:00",
        },
        xColumnNames: ["X_Value"],
        yColumnNames: [
          "Chamber Pressure",
          "Injector Temp",
          "Light Sensor",
          "Nitrous Bottle",
          "NOX Pressure",
          "Ox Tank Load Cell",
        ],
      },
    }}
  />
</main>
