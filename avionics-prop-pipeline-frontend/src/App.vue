<script setup lang="ts">
import BucketUploadForm from './components/BucketUploadForm.vue';
import BucketUploads from './components/BucketUploads.vue';
</script>

<template>
  <main>
    <h1>File Management</h1>
    <BucketUploadForm />
    <BucketUploads />
  </main>
</template>

<style scoped>
main {
  background: #e8f0fe;
  padding: 20px;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
h1 {
  margin-bottom: 0px;
  color: #333;
}
</style>
