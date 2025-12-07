<template>
  <div class="home">
    <h2>Welcome to 3focus</h2>
    <p>This is the home page.</p>

    <div class="api-status">
      <h3>Backend API Status</h3>
      <button @click="checkBackend">Check Backend</button>
      <p v-if="backendStatus" :class="backendStatus.success ? 'success' : 'error'">
        {{ backendStatus.message }}
      </p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Home',
  data() {
    return {
      backendStatus: null
    }
  },
  methods: {
    async checkBackend() {
      try {
        const response = await axios.get('http://localhost:8080/health')
        this.backendStatus = {
          success: true,
          message: response.data.message
        }
      } catch (error) {
        this.backendStatus = {
          success: false,
          message: 'Failed to connect to backend: ' + error.message
        }
      }
    }
  }
}
</script>

<style scoped>
.home {
  max-width: 800px;
  margin: 0 auto;
}

h2 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

.api-status {
  margin-top: 2rem;
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 8px;
}

button {
  background-color: #42b983;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  margin: 1rem 0;
}

button:hover {
  background-color: #359268;
}

.success {
  color: #42b983;
  font-weight: bold;
}

.error {
  color: #e74c3c;
  font-weight: bold;
}
</style>
