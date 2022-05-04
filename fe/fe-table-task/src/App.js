import "./App.css";
import { FilteringTable } from "./components/FilteringTable";
import { QueryClientProvider, QueryClient } from 'react-query'

function App() {
  const queryClient= new QueryClient();
  return (
    <QueryClientProvider client={queryClient}>
      < div className="App">
        <FilteringTable/>
      </div>
    </QueryClientProvider>

  );
}

export default App;
