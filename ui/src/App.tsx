import './App.sass';
import { Header } from './components/Header/Header';
import { StatusData } from './components/StatusData/StatusData';
import { ChangeLog } from './components/ChangeLog/ChangeLog';
import { PathList } from './components/PathList/PathList';
import { dataService } from './DataService';
import { useState } from 'react';
import { CollapsedContext, IAccordion } from './CollapsedContext';

const status = dataService.getStatus()
const changeLog = dataService.getData()
const totalChanges = dataService.getTotalChanged()

let test = (window as { [key: string]: any })["reportData"] as object;
console.log(test);

function App() {
  const [accordions, setAccordions] = useState([] as IAccordion[])
  const setCollapseAll = () => {
    setAccordions((prev) => {
      return prev.map(x => { return { ...x, isCollpased: true } })
    })
  }

  return (
    <div className="App">
      <Header dateGenerated={status?.startTime}></Header>
      <StatusData baseFile={status?.baseFile} secondFile={status?.secondFile} executionTime={status?.executionTime} totalPathChanges={totalChanges} flags={Object.keys(status?.executionFlags ?? {}).length}></StatusData>
      <ChangeLog onCollapseAll={setCollapseAll}></ChangeLog>
      <CollapsedContext.Provider value={{ accordions: accordions, setAccordions: setAccordions }}>
        <PathList changeList={changeLog}></PathList>
      </CollapsedContext.Provider>
    </div >
  );
}

export default App;
