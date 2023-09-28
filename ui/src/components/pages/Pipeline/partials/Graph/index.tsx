import React, {
  createContext,
  MouseEvent,
  useCallback,
  useEffect,
  useMemo,
  useState,
  useContext,
} from "react";

import ReactFlow, {
  applyEdgeChanges,
  applyNodeChanges,
  Edge,
  EdgeChange,
  Node,
  NodeChange,
  NodeTypes,
  EdgeTypes,
  Position,
  Panel,
  useReactFlow,
  useViewport,
  ReactFlowProvider,
} from "reactflow";
import Card from "@mui/material/Card";
import IconButton from "@mui/material/IconButton";
import { graphlib, layout } from "dagre";
import EdgeInfo from "./partials/EdgeInfo";
import NodeInfo from "./partials/NodeInfo";
import Spec from "./partials/Spec";
import CustomEdge from "./partials/CustomEdge";
import CustomNode from "./partials/CustomNode";
import { AppContext } from "../../../../../App";
import { AppContextProps } from "../../../../../types/declarations/app";
import { SidebarType } from "../../../../common/SlidingSidebar";
import {
  FlowProps,
  GraphProps,
  HighlightContextProps,
} from "../../../../../types/declarations/graph";
import lock from "../../../../../images/lock.svg";
import unlock from "../../../../../images/unlock.svg";
import scrollToggle from "../../../../../images/move-arrows.svg";
import closedHand from "../../../../../images/closed.svg";
import fullscreen from "../../../../../images/fullscreen.svg";
import sidePanel from "../../../../../images/side-panel.svg";
import zoomInIcon from "../../../../../images/zoom-in.svg";
import zoomOutIcon from "../../../../../images/zoom-out.svg";
import source from "../../../../../images/source.svg";
import map from "../../../../../images/map.svg";
import reduce from "../../../../../images/reduce.svg";
import sink from "../../../../../images/sink.svg";
import input from "../../../../../images/input.svg";
import generator from "../../../../../images/generator.svg";

import "reactflow/dist/style.css";
import "./style.css";

const nodeWidth = 252;
const nodeHeight = 72;
const graphDirection = "LR";

const defaultNodeTypes: NodeTypes = {
  custom: CustomNode,
};

const defaultEdgeTypes: EdgeTypes = {
  custom: CustomEdge,
};

//sets nodes and edges to highlight in the graph
export const HighlightContext = createContext<HighlightContextProps>(null);

const getLayoutedElements = (
  nodes: Node[],
  edges: Edge[],
  direction: string
) => {
  const dagreGraph = new graphlib.Graph();
  dagreGraph.setDefaultEdgeLabel(() => ({}));
  const isHorizontal = direction === "LR";
  dagreGraph.setGraph({ rankdir: direction, ranksep: 240, edgesep: 180 });

  nodes.forEach((node) => {
    dagreGraph.setNode(node.id, { width: nodeWidth, height: nodeHeight });
  });

  edges.forEach((edge) => {
    // skipping the side input edges to ensure graph alignment without the sideInputs
    if (!edge?.data?.sideInputEdge)
      dagreGraph.setEdge(edge.source, edge.target);
  });

  layout(dagreGraph);
  // setting nodes excepts generators and assigning generator height after remaining nodes are laid out
  let max_pos = 0;
  nodes.forEach((node) => {
    if (node?.data?.type !== "sideInput") {
      const nodeWithPosition = dagreGraph.node(node.id);
      max_pos = Math.max(max_pos, nodeWithPosition.y);
    }
  });
  let cnt = 2;
  nodes.forEach((node) => {
    const nodeWithPosition = dagreGraph.node(node.id);
    if (node?.data?.type === "sideInput") {
      nodeWithPosition.x = nodeWidth;
      nodeWithPosition.y = max_pos + nodeHeight * cnt;
      cnt++;
    }
    node.targetPosition = isHorizontal ? Position.Left : Position.Top;
    node.sourcePosition = isHorizontal ? Position.Right : Position.Bottom;

    // We are shifting the dagre node position (anchor=center) to the top left
    // so that it matches the React Flow node anchor point (top left).
    node.position = {
      x: nodeWithPosition.x - nodeWidth / 2,
      y: nodeWithPosition.y - nodeHeight / 2,
    };
  });

  return { nodes, edges };
};

const Flow = (props: FlowProps) => {
  const zoomLevel = useViewport().zoom;
  const { zoomIn, zoomOut, fitView } = useReactFlow();
  const [isLocked, setIsLocked] = useState(false);
  const [isPanOnScrollLocked, setIsPanOnScrollLocked] = useState(false);
  const {
    nodes,
    edges,
    onNodesChange,
    onEdgesChange,
    handleNodeClick,
    handleEdgeClick,
    handlePaneClick,
  } = props;

  const onIsLockedChange = useCallback(
    () => setIsLocked((prevState) => !prevState),
    []
  );
  const onIsPanOnScrollLockedChange = useCallback(
    () => setIsPanOnScrollLocked((prevState) => !prevState),
    []
  );
  const onFullScreen = useCallback(() => fitView(), [zoomLevel]);
  const onZoomIn = useCallback(() => zoomIn({ duration: 500 }), [zoomLevel]);
  const onZoomOut = useCallback(() => zoomOut({ duration: 500 }), [zoomLevel]);

  return (
    <ReactFlow
      nodeTypes={defaultNodeTypes}
      edgeTypes={defaultEdgeTypes}
      nodes={nodes}
      edges={edges}
      onNodesChange={onNodesChange}
      onEdgesChange={onEdgesChange}
      onEdgeClick={handleEdgeClick}
      onNodeClick={handleNodeClick}
      onPaneClick={handlePaneClick}
      fitView
      preventScrolling={!isLocked}
      panOnDrag={!isLocked}
      panOnScroll={isPanOnScrollLocked}
      maxZoom={2.75}
    >
      <Panel position="bottom-left" className={"interaction"}>
        <IconButton onClick={onIsLockedChange}>
          <img src={isLocked ? lock : unlock} alt={"lock-unlock"} />
        </IconButton>
        <IconButton onClick={onIsPanOnScrollLockedChange}>
          <img
            src={isPanOnScrollLocked ? closedHand : scrollToggle}
            alt={"panOnScrollLock"}
          />
        </IconButton>
        <div className={"divider"} />
        <IconButton onClick={onFullScreen}>
          <img src={fullscreen} alt={"fullscreen"} />
        </IconButton>
        <IconButton
          onClick={() => {
            //TODO add single panel logic
            alert("sidePanel");
          }}
        >
          <img src={sidePanel} alt={"sidePanel"} />
        </IconButton>
        <div className={"divider"} />
        <IconButton onClick={onZoomIn}>
          <img src={zoomInIcon} alt="zoom-in" />
        </IconButton>
        <IconButton onClick={onZoomOut}>
          <img src={zoomOutIcon} alt="zoom-out" />
        </IconButton>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="54"
          height="30"
          viewBox="0 0 54 30"
          fill="none"
        >
          <g>
            <rect
              x="0.5"
              y="0.5"
              width="53"
              height="29"
              rx="3.5"
              fill="white"
              stroke="#D4D7DC"
            />
            <text x="50%" y="50%" className={"zoom-percent-text"}>
              {(((zoomLevel - 0.5) / 1.5) * 100).toFixed(0)}%
            </text>
          </g>
        </svg>
      </Panel>
      <Panel position="bottom-left" className={"legend"}>
        <div className={"legend-title"}>
          <img src={source} alt={"source"} />
          <div className={"legend-text"}>Source</div>
        </div>
        <div className={"legend-title"}>
          <img src={map} alt={"map"} />
          <div className={"legend-text"}>Map</div>
        </div>
        <div className={"legend-title"}>
          <img src={reduce} alt={"reduce"} />
          <div className={"legend-text"}>Reduce</div>
        </div>
        <div className={"legend-title"}>
          <img src={sink} alt={"sink"} />
          <div className={"legend-text"}>Sink</div>
        </div>
        <div className={"legend-title"}>
          <img src={input} width={22} height={24} alt={"input"} />
          <div className={"legend-text"}>Input</div>
        </div>
        <div className={"legend-title"}>
          <img src={generator} width={22} height={24} alt={"generator"} />
          <div className={"legend-text"}>Generator</div>
        </div>
      </Panel>
    </ReactFlow>
  );
};

// hides the side input edges
const hide = (hidden: { [x: string]: any }) => (edge: Edge) => {
  edge.hidden = hidden[edge?.data?.source];
  return edge;
};

const getHiddenValue = (edges: Edge[]) => {
  const hiddenEdges = {};
  edges?.forEach((edge) => {
    if (edge?.data?.sideInputEdge) {
      hiddenEdges[edge?.data?.source] = true;
    }
  });
  return hiddenEdges;
};

export default function Graph(props: GraphProps) {
  const { data, namespaceId, pipelineId } = props;
  const { setSidebarProps } = useContext<AppContextProps>(AppContext);

  const { nodes: layoutedNodes, edges: layoutedEdges } = useMemo(() => {
    return getLayoutedElements(data.vertices, data.edges, graphDirection);
  }, [data]);

  const [nodes, setNodes] = useState<Node[]>(layoutedNodes);
  const [edges, setEdges] = useState<Edge[]>(layoutedEdges);
  const [sideNodes, setSideNodes] = useState<Map<string, Node>>(new Map());
  const [sideEdges, setSideEdges] = useState<Map<string, string>>(new Map());
  const initialHiddenValue = getHiddenValue(layoutedEdges);
  const [hidden, setHidden] = useState(initialHiddenValue);

  useEffect(() => {
    const nodeSet = new Map();
    layoutedNodes.forEach((node) => {
      if (node?.data?.sideHandle) {
        nodeSet.set(node?.data?.name, node);
      }
    });
    setSideNodes(nodeSet);
  }, [layoutedNodes]);

  useEffect(() => {
    const edgeSet: Map<string, string> = new Map();
    layoutedEdges.forEach((edge) => {
      if (edge?.data?.sideInputEdge) {
        edgeSet.set(edge?.id, edge?.targetHandle);
      }
    });
    setSideEdges(edgeSet);
  }, [layoutedEdges]);

  useEffect(() => {
    setNodes(layoutedNodes);
    setEdges(layoutedEdges);
  }, [data]);

  const onNodesChange = useCallback(
    (changes: NodeChange[]) =>
      setNodes((nds) => applyNodeChanges(changes, nds)),
    [setNodes]
  );
  const onEdgesChange = useCallback(
    (changes: EdgeChange[]) =>
      setEdges((eds) => applyEdgeChanges(changes, eds)),
    [setEdges]
  );

  const [edgeOpen, setEdgeOpen] = useState(false);

  const [edgeId, setEdgeId] = useState<string>();
  const [edge, setEdge] = useState<Edge>();

  const handleEdgeClick = useCallback(
    (event: MouseEvent, edge: Edge) => {
      setEdge(edge);
      setEdgeId(edge.id);
      setEdgeOpen(true);
      setShowSpec(false);
      setNodeOpen(false);
      if (setSidebarProps) {
        setSidebarProps({
          type: SidebarType.EDGE_DETAILS,
          edgeDetailsProps: {
            namespaceId,
            edgeId: edge.id,
          },
        });
      }
    },
    [setSidebarProps, namespaceId]
  );

  // This has been added to make sure that edge container refreshes on edges being refreshed
  useEffect(() => {
    let flag = false;
    edges.forEach((dataEdge) => {
      const edge = {} as Edge;
      if (dataEdge.id === edgeId && !flag) {
        edge.data = dataEdge.data;
        edge.label = dataEdge.label;
        edge.id = dataEdge.id;
        flag = true;
        setEdge(edge);
      }
    });
  }, [edges, edgeId]);

  const [nodeOpen, setNodeOpen] = useState(false);

  const [nodeId, setNodeId] = useState<string>();
  const [node, setNode] = useState<Node>();

  useEffect(() => {
    setEdges((eds) => eds.map(hide(hidden)));
  }, [hidden, data]);

  const handleNodeClick = useCallback(
    (event: MouseEvent, node: Node) => {
      setNode(node);
      setNodeId(node.id);
      setNodeOpen(true);
      setShowSpec(false);
      setEdgeOpen(false);
      setHidden((prevState) => {
        const updatedState = {};
        Object.keys(prevState).forEach((key) => {
          updatedState[key] =
            node?.data?.type === "sideInput" ? !(key === node.id) : true;
        });
        return updatedState;
      });
      if (setSidebarProps) {
        if (node?.data?.type === "sideInput") {
          setSidebarProps({
            type: SidebarType.GENERATOR_DETAILS,
            generatorDetailsProps: {
              namespaceId,
              pipelineId,
              vertexId: node.id,
              generatorDetails: sideNodes.get(node.id) || {},
            },
          });
        } else {
          setSidebarProps({
            type: SidebarType.VERTEX_DETAILS,
            vertexDetailsProps: {
              namespaceId,
              pipelineId,
              vertexId: node.id,
              vertexSpecs: data.pipeline?.spec?.vertices || [],
            },
          });
        }
      }
    },
    [
      setSidebarProps,
      namespaceId,
      pipelineId,
      data.pipeline,
      setHidden,
      sideNodes,
    ]
  );

  // This has been added to make sure that node container refreshes on nodes being refreshed
  useEffect(() => {
    nodes.forEach((dataNode) => {
      const node = {} as Node;
      if (dataNode.id === nodeId) {
        node.data = dataNode.data;
        node.id = dataNode.id;
        setNode(node);
      }
    });
  }, [nodes, nodeId]);

  const [highlightValues, setHighlightValues] = useState<{
    [key: string]: boolean;
  }>({});

  const handlePaneClick = () => {
    setShowSpec(true);
    setEdgeOpen(false);
    setNodeOpen(false);
    setHighlightValues({});
    setHidden((prevState) => {
      const updatedState = {};
      Object.keys(prevState).forEach((key) => {
        updatedState[key] = true;
      });
      return updatedState;
    });
  };

  const [showSpec, setShowSpec] = useState(true);

  return (
    <div style={{ height: "90%" }}>
      <div className="Graph" data-testid="graph">
        <HighlightContext.Provider
          value={{
            highlightValues,
            setHighlightValues,
            setHidden,
            handleNodeClick,
            sideInputNodes: sideNodes,
            sideInputEdges: sideEdges,
          }}
        >
          <ReactFlowProvider>
            <Flow
              {...{
                nodes,
                edges,
                onNodesChange,
                onEdgesChange,
                handleNodeClick,
                handleEdgeClick,
                handlePaneClick,
              }}
            />
          </ReactFlowProvider>
        </HighlightContext.Provider>
      </div>

      {/*<Card*/}
      {/*  sx={{ borderBottom: 1, borderColor: "divider", boxShadow: 1 }}*/}
      {/*  data-testid={"card"}*/}
      {/*  variant={"outlined"}*/}
      {/*>*/}
      {/*  {showSpec && <Spec pipeline={data.pipeline} />}*/}
      {/*  {edgeOpen && <EdgeInfo data-testid="edge-info" edge={edge} />}*/}
      {/*  {nodeOpen && (*/}
      {/*    <NodeInfo*/}
      {/*      data-testid="node-info"*/}
      {/*      node={node}*/}
      {/*      namespaceId={namespaceId}*/}
      {/*      pipelineId={pipelineId}*/}
      {/*    />*/}
      {/*  )}*/}
      {/*</Card>*/}
    </div>
  );
}
