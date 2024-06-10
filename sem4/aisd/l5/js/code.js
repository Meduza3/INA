console.log("Hello!")

function drawCompleteGraph(n) {
  const width = 800;
  const height = 800;
  const radius = 300; // Radius of the circle on which nodes will be placed
  const svg = d3.select("#graph").append("svg")
      .attr("width", width)
      .attr("height", height);

  // Generate nodes data
  const nodes = Array.from({ length: n }, (_, i) => ({
      id: i,
      x: width / 2 + radius * Math.cos(2 * Math.PI * i / n),
      y: height / 2 + radius * Math.sin(2 * Math.PI * i / n)
  }));

  // Generate links data (all possible connections between nodes)
  const links = [];
  for (let i = 0; i < n; i++) {
      for (let j = i + 1; j < n; j++) {
          links.push({ source: i, target: j });
      }
  }

  // Draw links (edges)
  svg.selectAll(".link")
      .data(links)
      .enter().append("line")
      .attr("class", "link")
      .attr("x1", d => nodes[d.source].x)
      .attr("y1", d => nodes[d.source].y)
      .attr("x2", d => nodes[d.target].x)
      .attr("y2", d => nodes[d.target].y)
      .style("stroke", "#aaa");

  // Draw nodes
  svg.selectAll(".node")
      .data(nodes)
      .enter().append("circle")
      .attr("class", "node")
      .attr("cx", d => d.x)
      .attr("cy", d => d.y)
      .attr("r", 10)
      .style("fill", "red");
}

drawCompleteGraph(10); // Call this function with any n value
