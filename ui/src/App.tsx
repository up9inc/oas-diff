import './App.sass';
import { StatusData } from './components/StatusData/StatusData';
import PathList from './components/PathList/PathList';
import { dataService } from './DataService';
import { Header } from './components/Header/Header';

const status = dataService.getStatus()
const changeLog = dataService.getData()
const totalChanges = dataService.getTotalChanged()

const App = () => {
  return (
    <div className="App">
      <Header dateGenerated={status?.startTime} />
      <StatusData baseFile={status?.baseFile} secondFile={status?.secondFile} executionTime={status?.executionTime} totalPathChanges={totalChanges} flags={Object.keys(status?.executionFlags ?? {}).length} />
      <PathList changeList={changeLog} />
    </div >
  );
}

export default App;
