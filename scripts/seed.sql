TRUNCATE missions, personnel, assets CASCADE;

-- Missions
INSERT INTO missions (name, description, status, mission_type, start_time, end_time) VALUES
  ('Operation Iron Falcon',   'Reconnaissance sweep of sector 7',         'active',    'reconnaissance', '2026-07-01 06:00:00+00', '2026-07-01 18:00:00+00'),
  ('Operation Silent Strike', 'Direct action against high-value target',  'planned',   'direct_action',  '2026-07-15 02:00:00+00', '2026-07-15 06:00:00+00'),
  ('Operation Desert Watch',  'Surveillance of forward supply routes',    'completed', 'surveillance',   '2026-06-01 00:00:00+00', '2026-06-10 00:00:00+00');

-- Personnel
INSERT INTO personnel (rank, first_name, last_name, unit_designator, clearance_level, status) VALUES
  ('Captain',            'James',   'Mercer',   '1st Special Forces Group',  'top_secret', 'active'),
  ('Staff Sergeant',     'Maria',   'Reyes',    '75th Ranger Regiment',      'secret',     'active'),
  ('Lieutenant Colonel', 'David',   'Okafor',   'JSOC',                      'top_secret', 'active'),
  ('Sergeant',           'Emily',   'Tran',     '82nd Airborne Division',    'secret',     'inactive'),
  ('Major',              'Carlos',  'Herrera',  '160th SOAR',                'top_secret', 'active');

-- Assets
INSERT INTO assets (designation, asset_type, status, notes) VALUES
  ('AH-64 Apache #7',   'rotary_wing',  'available', 'Fully armed and fueled'),
  ('MQ-9 Reaper #3',    'drone',        'deployed',  'On station over sector 7'),
  ('M1A2 Abrams #12',   'armor',        'available', 'Requires track maintenance'),
  ('C-130J #5',         'fixed_wing',   'available', 'Configured for HALO insertion'),
  ('RQ-11 Raven #9',    'drone',        'available', 'Portable ISR platform');
