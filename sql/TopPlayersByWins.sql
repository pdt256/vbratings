-- Top Teams by Total Wins
SELECT
  PlayerA.name
  ,PlayerB.name
  ,SUM(CASE WHEN Match.teamA_id = Team.id THEN 1 ELSE 0 END) AS Wins
  ,SUM(CASE WHEN Match.teamB_id = Team.id THEN 1 ELSE 0 END) AS Losses
  ,SUM(CASE WHEN Match.teamA_id = Team.id AND SetScore.teamBForfeit THEN 1 ELSE 0 END) AS WinsByForfeit
FROM
  Team
  INNER JOIN Player AS PlayerA ON PlayerA.id = Team.playerA_id
  INNER JOIN Player AS PlayerB ON PlayerB.id = Team.playerB_id
  LEFT JOIN Match ON (Match.teamA_id = Team.id OR Match.teamB_id = Team.id)
  LEFT JOIN SetScore ON SetScore.match_id = Match.id
GROUP BY Team.id
ORDER BY Wins DESC
;

-- Top Teams by Win %
SELECT
  PlayerA.name
  ,PlayerB.name
  ,SUM(CASE WHEN Match.teamA_id = Team.id THEN 1 ELSE 0 END) AS Wins
  ,SUM(CASE WHEN Match.teamB_id = Team.id THEN 1 ELSE 0 END) AS Losses
  ,ROUND(SUM(CASE WHEN Match.teamA_id = Team.id THEN 1 ELSE 0 END) * 100.0 / COUNT(Match.id), 2) AS WinPct
FROM
  Team
  INNER JOIN Player AS PlayerA ON PlayerA.id = Team.playerA_id
  INNER JOIN Player AS PlayerB ON PlayerB.id = Team.playerB_id
  LEFT JOIN Match ON (Match.teamA_id = Team.id OR Match.teamB_id = Team.id)
  LEFT JOIN SetScore ON SetScore.match_id = Match.id
GROUP BY Team.id
ORDER BY WinPct DESC
;
