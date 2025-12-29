<script>
  // Esta importación la genera Wails automáticamente al compilar
  import {AnalyzeDomain} from '../wailsjs/go/main/App';

  let domain = "";
  let result = null;
  let loading = false;
  let errorMsg = "";
  let activeTab = 'analizar'; // Control de pestañas

  async function startScan() {
    if (!domain) return;
    loading = true;
    errorMsg = "";
    result = null;

    try {
      // Llamada directa a tu función de Go
      result = await AnalyzeDomain(domain);
    } catch (err) {
      errorMsg = "Error: " + err;
    } finally {
      loading = false;
    }
  }
</script>

<main>
  <nav>
    <button class:active={activeTab === 'analizar'} on:click={() => activeTab = 'analizar'}>🔍 Analizador</button>
    <button class:active={activeTab === 'info'} on:click={() => activeTab = 'info'}>📚 ¿Qué es SSL?</button>
  </nav>

  {#if activeTab === 'analizar'}
    <div class="content">
      <h1>Verificador SSL Labs</h1>
      <div class="input-group">
        <input bind:value={domain} placeholder="google.com" disabled={loading} type="text"/>
        <button on:click={startScan} disabled={loading}>
          {loading ? 'Analizando (Espere...)' : 'Analizar'}
        </button>
      </div>

      {#if loading}
        <p class="pulse">Contactando con SSL Labs...</p>
      {/if}

      {#if errorMsg}
        <div class="alert error">{errorMsg}</div>
      {/if}

      {#if result}
        <div class="report">
          <h2>Resultados para: {result.host}</h2>
          <span class="badge {result.status === 'READY' ? 'success' : 'warn'}">{result.status}</span>
          
          <table class="results-table">
            <thead>
              <tr>
                <th>IP Servidor</th>
                <th>Grado</th>
                <th>Estado</th>
              </tr>
            </thead>
            <tbody>
              {#each result.endpoints as ep}
                <tr>
                  <td>{ep.ipAddress}</td>
                  <td class="grade {ep.grade?.startsWith('A') ? 'grade-a' : 'grade-bad'}">
                    {ep.grade || '-'}
                  </td>
                  <td>{ep.statusMessage}</td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      {/if}
    </div>
  {/if}

  {#if activeTab === 'info'}
    <div class="content info-section">
      <h2>¿Qué es SSL/TLS?</h2>
      
      

      <div class="card">
        <h3>🔐 SSL (Secure Sockets Layer)</h3>
        <p>Es la tecnología estándar antigua para mantener segura una conexión a internet. Protege los datos confidenciales que se envían entre dos sistemas.</p>
      </div>

      <div class="card">
        <h3>🛡️ TLS (Transport Layer Security)</h3>
        <p>Es la versión <strong>moderna y segura</strong> de SSL. Aunque seguimos diciendo "Certificado SSL", en realidad hoy usamos TLS.</p>
      </div>

      <h3>¿Qué significan las calificaciones?</h3>
      <ul>
        <li><strong style="color:green">A+ / A:</strong> Configuración excelente. Sitio seguro.</li>
        <li><strong style="color:orange">B / C:</strong> Configuración obsoleta. Soporta protocolos viejos (TLS 1.0/1.1).</li>
        <li><strong style="color:red">F:</strong> Peligroso. Vulnerable a ataques conocidos.</li>
      </ul>
    </div>
  {/if}
</main>

<style>
  /* Estilos básicos para que se vea profesional */
  main { font-family: sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
  nav { display: flex; gap: 10px; margin-bottom: 20px; border-bottom: 2px solid #eee; }
  nav button { background: none; border: none; padding: 10px 20px; cursor: pointer; font-size: 16px; opacity: 0.6; }
  nav button.active { border-bottom: 3px solid #007bff; opacity: 1; font-weight: bold; }
  
  .input-group { display: flex; gap: 10px; margin-bottom: 20px; }
  input { flex: 1; padding: 10px; border: 1px solid #ccc; border-radius: 4px; }
  button { padding: 10px 20px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; }
  button:disabled { background: #ccc; }
  
  .results-table { width: 100%; border-collapse: collapse; margin-top: 15px; }
  th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
  
  .grade { font-weight: bold; font-size: 1.2em; }
  .grade-a { color: #28a745; }
  .grade-bad { color: #dc3545; }
  
  .info-section .card { background: #f9f9f9; padding: 15px; margin-bottom: 15px; border-left: 4px solid #007bff;color: #000000; }
  .pulse { animation: pulse 1.5s infinite; color: #666; }
  @keyframes pulse { 0% { opacity: 0.5; } 50% { opacity: 1; } 100% { opacity: 0.5; } }
</style>
