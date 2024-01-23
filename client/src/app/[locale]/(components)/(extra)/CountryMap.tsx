import React, { useEffect } from "react";
import "leaflet/dist/leaflet.css";
import L from "leaflet";

interface CountryMapProps {
  width?: string;
  height?: string;
}

const CountryMap: React.FC<CountryMapProps> = ({
  width = "100%",
  height = "100%",
}) => {
  useEffect(() => {
    // Create a Leaflet map
    const map = L.map("map").setView([51.505, -0.09], 13);

    // Add a tile layer to the map
    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution: "© OpenStreetMap contributors",
    }).addTo(map);
  }, []); // Empty dependency array ensures useEffect runs once

  return <div id="map" style={{ width, height }} />;
};

export default CountryMap;
