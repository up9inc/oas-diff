import './App.sass';
import { useState } from 'react'
import { StatusData } from './components/StatusData/StatusData';
import { ChangeLog } from './components/ChangeLog/ChangeLog';
import PathList from './components/PathList/PathList';
import { dataService } from './DataService';
import { CollapsedContext, IAccordion } from './CollapsedContext';
import { Header } from './components/Header/Header';
import { RecoilRoot } from 'recoil';

const status = dataService.getStatus()
const changeLog = dataService.getData()
const totalChanges = dataService.getTotalChanged()

const App = () => {
  return (
    <div className="App">
      <Header dateGenerated={status?.startTime} />
      <StatusData baseFile={status?.baseFile} secondFile={status?.secondFile} executionTime={status?.executionTime} totalPathChanges={totalChanges} flags={Object.keys(status?.executionFlags ?? {}).length} />
      <RecoilRoot>
        <ChangeLog />
        <PathList changeList={changeLog} />
      </RecoilRoot>
    </div >
  );
}

export default App;
