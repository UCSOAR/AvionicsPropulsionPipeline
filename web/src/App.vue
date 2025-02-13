<script setup lang="ts">
import { ref, provide } from 'vue';
import BucketUploadForm from './components/BucketUploadForm.vue';
import BucketUploads from './components/BucketUploads.vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faSun, faMoon } from '@fortawesome/free-solid-svg-icons';
import ChartPlaceholder from './components/ChartPlaceholder.vue';
import ColumnTray from './components/ColumnTray.vue';

const isDarkMode = ref(true);

// Provide the state to child components
provide('isDarkMode', isDarkMode);

const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value;
  const root = document.documentElement;
  root.classList.toggle('dark', isDarkMode.value);
};
</script>

<template>
  <main :class="{ dark: isDarkMode }">
    <header>
      <div>
        <!-- Conditional Rendering for Light/Dark Mode Logo -->
        <img
          v-if="isDarkMode"
          src="/Logo/soar-logo-dark.svg"
          alt="Soar Logo"
          class="logo"
        />
        <img
          v-else
          src="/Logo/soar-logo-light.svg"
          alt="Soar Logo"
          class="logo"
        />
      </div>
      <button @click="toggleTheme" class="theme-toggle-button">
        <FontAwesomeIcon :icon="isDarkMode ? faSun : faMoon" />
      </button>
    </header>

    <!-- Content wrapper for alignment -->
    <div class="content-wrapper">
      <div class="files">
        <BucketUploadForm />
        <BucketUploads />
      </div>

      <!-- Chart Component -->
      <div class="graph">
        <ColumnTray />
        <ChartPlaceholder class="chart-container"/>
      </div>
    </div>
  </main>
</template>

<style>
/* Reset Styles */
html, body {
  margin: 0 !important;
  padding: 0 !important;
  border: 0;
  box-sizing: border-box;
  width: 100%;
  height: 100%;
  font-family: 'Poppins', sans-serif;
  /* Hide scrolling while preserving the layout */
  overflow: hidden;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius:10px;
  
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 100px;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #555;
  border-radius:10px;
} 


/* Main Layout */
main {
  padding: 20px;
  height: 100vh; /* Changed from min-height: 100vh to height: 100vh */
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: background-color 0.3s ease, color 0.3s ease;
  box-sizing: border-box; /* Ensures padding is included within 100vh */
}

/* Header Styling */
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  margin-bottom: 40px;
}

.logo {
  width: 75px;
}

/* Content Wrapper */
.content-wrapper {
  min-height: 930px;
  display: flex;
  justify-content: center;
  align-self: stretch;
  width: 100%;
  gap: 10px; /* Reduced gap to bring components closer */
}

/* File Tray (Left Panel) */
.files {
  padding: 1px;
  display: flex;
  flex-direction: column;
  align-items: center;
  /* width: 45%;  */
}

.graph {
  padding-top: 14px;
  gap: 1px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  flex: 1;
  /* width: 45%;  */
}

/* Chart Container (Right Panel) */
.chart-container {
  width: 50%; /* Reduced width for alignment */
  height: 50%; /* Halved the height */
  min-height: 200px; /* Ensures responsiveness */
  display: flex;
  justify-content: flex-start;
  align-items: center;
  background: #1e1e1e; /* Background for visualization */
  border-radius: 8px;
  padding: 10px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.3);
}

/* Theme Toggle Button */
.theme-toggle-button {
  background: #007bff;
  color: white;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  font-weight: bold;
  transition: all 0.3s ease;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.theme-toggle-button:hover {
  background: #0056b3;
  transform: scale(1.05);
}

/* Dark Mode Styles */
main.dark {
  background-color: #121212;
  color: #ffffff;
}

.theme-toggle-button.dark {
  background: #ffffff;
  color: #333;
}

.theme-toggle-button.dark:hover {
  background: #dddddd;
}



</style>