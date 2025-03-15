<script lang="ts">
  import "$lib/styles/global.scss";
  import "$lib/components/IconButton.svelte";
  import Dashboard from "$lib/components/Dashboard.svelte";
</script>

<main>
  <Dashboard
    metadata={{
      operator: "Test Operator",
      resultTimestamp: {
        date: "2021-09-01",
        time: "12:00:00",
      },
      xColumnNames: ["x1", "x2", "x3"],
      yColumnNames: ["y1", "y2", "y3"],
    }}
  />
</main>
