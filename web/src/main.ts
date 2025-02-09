import { createApp } from "vue";
import App from "./App.vue";
import VueApexCharts from 'vue3-apexcharts';
import { createPinia } from 'pinia';


const app = createApp(App);
app.use(VueApexCharts);

const pinia = createPinia();
app.use(pinia);
app.component('apexchart', VueApexCharts);
app.mount('#app');  