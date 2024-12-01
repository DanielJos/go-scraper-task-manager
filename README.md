# Web Scraping Task Manager

This program will:
1. Accept tasks (URLs to scrape) via an API or a command-line interface.
2. Scrape the URLs concurrently, processing multiple tasks in parallel.
3. Maintain a status for each task (pending, in progress, completed, failed).
4. Collect and display statistics like:
	- Total tasks processed.
	- Average task runtime.
	- Success vs failure rate.
5. Allow you to query the status of tasks and view results.

## Features:
1. **Concurrency**:
	- Use Goroutines to handle scraping tasks concurrently.
	- Implement a worker pool to limit the number of simultaneous scrapes.
2. **Task Management**:
	- Use a simple in-memory database (e.g., a map) to track task states and results.
	- Optionally store results to a file or SQLite database.
3. **Statistics**:
	- Track start and completion times for tasks.
	- Calculate average runtimes and success rates dynamically.
4. **Extendability**:
	- Provide hooks for scraping logic to extract custom data (e.g., titles, meta tags, or other structured content).

## Architecture

### Layered
- **Handler** - Handle request parsing and returning
- **Service** - Business logic, handles connection to the data layer processes
- **Data** - Is the data store and the job manager

## Modules
- **Job Handler**
- **Data Handler**
- **Job Manager**

## Thoughts
- We'll need a factory pattern somewhere to spawn jobs, this will be the **job manager**.
- We'll need an API structure as well to handle the requests coming in. Data layer is the job manager, but lets keep this simple, maybe we just skip the service layer. I'll put a pin in that.

### Resources
1. Jobs (These do a task and save their results to **Data**)
2. Data (These represent the scraped data)

The **Jobs** resource will need to be able to do the following:
- Be created
- Execute
- Store it's results in the **Data** store
- It would be nice to be able to offload jobs to other processes, as well as handle concurrency within this runtime.

The **Data** resource will be the results, we'll need to be able to:
- Return individual data points
- Generate summary stats over the aggregation of the Data (over the Data store)

