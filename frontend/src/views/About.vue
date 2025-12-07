<template>
  <div class="about">
    <h2>About 3focus</h2>
    <p>A project built with Go backend and Vue frontend.</p>

    <div class="tech-stack">
      <h3>Technology Stack</h3>
      <ul>
        <li><strong>Backend:</strong> Go with Gin framework</li>
        <li><strong>Frontend:</strong> Vue 3 with Vite</li>
        <li><strong>Database:</strong> PostgreSQL</li>
        <li><strong>Development:</strong> Docker Compose</li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'About'
}
</script>

<style scoped>
.about {
  max-width: 800px;
  margin: 0 auto;
}

h2 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

.tech-stack {
  margin-top: 2rem;
  padding: 1rem;
  background-color: #f8f9fa;
  border-radius: 8px;
}

.tech-stack h3 {
  margin-bottom: 1rem;
  color: #2c3e50;
}

.tech-stack ul {
  list-style-position: inside;
}

.tech-stack li {
  padding: 0.5rem 0;
  border-bottom: 1px solid #e0e0e0;
}

.tech-stack li:last-child {
  border-bottom: none;
}
</style>
