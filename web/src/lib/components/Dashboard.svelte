<script lang="ts">
</script>

<div class="container">
  <h1>Dashboard</h1>
  <div class="content">
    <p>Select a file from the sidebar to view its data.</p>
  </div>
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.container {
    padding: 1rem;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    flex-grow: 1;

    h1 {
      margin: 0;
    }

    div.content {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-grow: 1;
      text-align: center;
    }
  }
</style>
