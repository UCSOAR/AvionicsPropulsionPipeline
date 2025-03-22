<script lang="ts">
  import "$lib/styles/global.scss";
  import "$lib/components/IconButton.svelte";
  import Dashboard from "$lib/components/Dashboard.svelte";
  import TopBar from "$lib/components/TopBar.svelte";
  import Sidebar from "$lib/components/Sidebar.svelte";
</script>

<main>
  <TopBar />
  <Sidebar />
  <Dashboard

    selectedFile={{
      name: "RHD324u9",
      metadata: {
        operator: "Test Operator",
        resultTimestamp: {
          date: "2021-09-01",
          time: "12:00:00",
        },
        xColumnNames: ["ewffgtrg", "x2", "x3"],
        yColumnNames: ["rthtrthrhtr", "y2", "y3"],
      },
    }}
  />
</main>

