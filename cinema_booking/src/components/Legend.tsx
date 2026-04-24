const Legend = () => (
  <div className="legend">
    <div className="legend-item"><div className="legend-swatch" style={{ background: 'var(--available)' }}></div>Available</div>
    <div className="legend-item"><div className="legend-swatch" style={{ background: 'var(--held-mine)' }}></div>Your hold</div>
    <div className="legend-item"><div className="legend-swatch" style={{ background: 'var(--held-other)' }}></div>Other hold</div>
    <div className="legend-item"><div className="legend-swatch" style={{ background: 'var(--confirmed)' }}></div>Confirmed</div>
  </div>
);


export default Legend