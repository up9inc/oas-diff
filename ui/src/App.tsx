import './App.sass';
import { Header } from './components/Header/Header';
import { GeneralData } from './components/GeneralData/GeneralData';
import { ChangeLog } from './components/ChangeLog/ChangeLog';
import { PathList } from './components/PathList/PathList';



function App() {
  return (
    <div className="App">
      <Header dateGenerated=''></Header>
      <GeneralData baseFile={'asdasd'} secondFile={'asdasd'} executionTime={'asdasd'} totalPathChanges={0} flags={0}></GeneralData>
      <ChangeLog></ChangeLog>
      <PathList></PathList>
      {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <div className='reportContainer'>
        <Card>
          <CardActionArea>
            <CardContent>
              <div className='details'>
                <span className='detailItem'>
                  <span className='itemTitle'>Report generated on: <br /></span>
                  <span className='itemData'>May 10 17:21:08.311</span>
                </span>
                <span className='detailItem'>
                  <span className='itemTitle'>Base File: <br /></span>
                  <span className='itemData'>/Users/haiut/Work/OAS/1/u1.json</span>
                </span>
                <span className='detailItem'>
                  <span className='itemTitle'>Second File: <br /></span>
                  <span className='itemData'>/Users/haiut/Work/OAS/1/u2.json</span>
                </span>
                <span className='detailItem'>
                  <span className='itemTitle'>Execution Time<br /></span>
                  <span className='itemData'>	261.672333ms</span>
                </span>
                <span className='detailItem'>
                  <span className='itemTitle'>Second File: <br /></span>
                  <span className='itemData'>2.json</span>
                </span>
              </div>
            </CardContent>
          </CardActionArea>
        </Card>

        <div className='filters'>

          <Card>
            <CardActionArea>
              <CardContent>
                <h1>Filters</h1>
                <FormControl size="small" sx={{ m: 1, minWidth: 150 }}>
                  <InputLabel id="demo-simple-select-label">Change Type</InputLabel>
                  <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    label="Change Type"
                  >
                    <MenuItem value={10}>Ten</MenuItem>
                    <MenuItem value={20}>Twenty</MenuItem>
                    <MenuItem value={30}>Thirty</MenuItem>
                  </Select>
                </FormControl>
                <FormControl size="small" sx={{ m: 1, minWidth: 120 }}>
                  <InputLabel id="demo-simple-select-label">Services</InputLabel>
                  <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    label="Services"
                  >
                    <MenuItem value={10}>Ten</MenuItem>
                    <MenuItem value={20}>Twenty</MenuItem>
                    <MenuItem value={30}>Thirty</MenuItem>
                  </Select>
                </FormControl>
              </CardContent>
            </CardActionArea>
          </Card>
        </div>
        <div className='report'>

          {list.map((item, index) => {
            return (
              <Accordion>
                <AccordionSummary
                  expandIcon={<ExpandMoreIcon />}
                  aria-controls="panel2a-content"
                  id="panel2a-header"
                >
                  <Typography>[PATHS] /customers/&#123;customerId&#125;addresses</Typography>
                </AccordionSummary>
                <AccordionDetails>
                  <Typography>
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
                    malesuada lacus ex, sit amet blandit leo lobortis eget.
                  </Typography>
                </AccordionDetails>
              </Accordion>
            )
          })}

        </div>
      </div> */}
    </div >
  );
}

export default App;
